
---

# go/README.md — API/WS Gateway

```markdown
# Go API / WebSocket Gateway

Mirrors MQTT → HTTP/WS for UIs and provides a simple control plane.

- Subscribes to:
  - `grid/v1/status/sim` → caches SimStatus
  - `grid/v1/node/+` → caches NodeDelta and fans out over WebSocket
- Serves:
  - `GET /api/status` → `{tick, topo_hash, state}`
  - `GET /api/nodes/{id}` → last known `{node_id, v_pu, tick}`
  - `WS /ws` → push deltas; client can filter nodes:
    - send: `{"subscribe":{"nodes":["bus-1","bus-2"]}}`

## Layout

```
go/
├─ cmd/api/main.go # wire config + MQTT bridge + HTTP/WS
├─ pkg/
│ ├─ api/ # HTTP/WS handlers
│ ├─ cache/ # in-memory state
│ ├─ config/ # env loader
│ ├─ control/ # helpers to publish Command (future)
│ ├─ gridbus/ # MQTT wrapper + topic helpers
│ └─ pb/grid/v1/ # generated protobuf bindings
```


## Setup

Initialize module (from `go/`):

```bash
go mod init <your-module>/go
go get github.com/eclipse/paho.mqtt.golang@v1.4.3
go get github.com/gorilla/websocket@v1.5.1
go get google.golang.org/protobuf@v1.34.1
```

