package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	grid "yourmod/go/pkg/pb/grid/v1"
)

// --- WebSocket hub -----------------------------------------------------------

type WSUpdate struct {
	Node *grid.NodeDelta `json:"-"`
	// You can add Edge, Status later if you want to push those too.
}

type client struct {
	id        string
	conn      *websocket.Conn
	interests map[string]struct{} // nodeIds this client wants
	sendCh    chan []byte
	hub       *Hub
}

type Hub struct {
	mu        sync.RWMutex
	clients   map[*client]struct{}
	broadcast chan WSUpdate
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*client]struct{}),
		broadcast: make(chan WSUpdate, 1024),
	}
}

// Push from outside (e.g., cache when a NodeDelta is updated)
func (h *Hub) PushNodeDelta(nd *grid.NodeDelta) {
	select {
	case h.broadcast <- WSUpdate{Node: nd}:
	default:
		// Drop on overflow to avoid blocking the sim path.
	}
}

func (h *Hub) run() {
	for upd := range h.broadcast {
		if upd.Node != nil {
			// Encode a small JSON for browsers (you can stream protobuf binary if you prefer).
			msg := map[string]any{
				"type":    "node",
				"node_id": upd.Node.GetNodeId(),
				"v_pu":    upd.Node.GetVPu(), // optional field; 0 if unset
				"tick":    upd.Node.GetEnv().GetClock().GetTick(),
			}
			payload, _ := json.Marshal(msg)

			h.mu.RLock()
			for c := range h.clients {
				// Filter by interest
				if _, ok := c.interests[upd.Node.GetNodeId()]; ok || len(c.interests) == 0 {
					select {
					case c.sendCh <- payload:
					default:
						// If the client is slow, drop the message.
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

// --- HTTP handler ------------------------------------------------------------

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	// TODO: tighten origin checks if exposed publicly
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WSDeps interface {
	// Optional hooks if you want to query cache on subscribe requests, etc.
}

func (h *Hub) WSHandler(_ WSDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "upgrade failed", http.StatusBadRequest)
			return
		}
		c := &client{
			id:        r.RemoteAddr,
			conn:      conn,
			interests: make(map[string]struct{}),
			sendCh:    make(chan []byte, 256),
			hub:       h,
		}

		h.mu.Lock()
		h.clients[c] = struct{}{}
		h.mu.Unlock()

		go c.writer()
		go c.reader()
	}
}

func (c *client) reader() {
	defer c.close()

	// Client can send a JSON control message: {"subscribe":{"nodes":["bus-1","bus-2"]}}
	type subMsg struct {
		Subscribe struct {
			Nodes []string `json:"nodes"`
		} `json:"subscribe"`
	}
	_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		var m subMsg
		if err := json.Unmarshal(data, &m); err == nil && len(m.Subscribe.Nodes) > 0 {
			newInt := make(map[string]struct{}, len(m.Subscribe.Nodes))
			for _, id := range m.Subscribe.Nodes {
				newInt[id] = struct{}{}
			}
			c.interests = newInt
		}
	}
}

func (c *client) writer() {
	ticker := time.NewTicker(30 * time.Second) // WS ping
	defer ticker.Stop()
	for {
		select {
		case msg, ok := <-c.sendCh:
			if !ok {
				return
			}
			_ = c.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *client) close() {
	c.hub.mu.Lock()
	delete(c.hub.clients, c)
	c.hub.mu.Unlock()
	close(c.sendCh)
	_ = c.conn.Close()
}

// Optional helper if you want to run the hub inside main()
func RunHub(ctx context.Context, h *Hub) {
	go h.run()
	<-ctx.Done()
	log.Println("ws hub stopping")
}
