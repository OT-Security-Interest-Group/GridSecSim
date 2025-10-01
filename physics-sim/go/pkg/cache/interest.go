package cache

import "sync"

// Keeps per-WS-client interest sets (optional helper if you want hub-independent storage).

type Interests struct {
	mu   sync.RWMutex
	data map[string]map[string]struct{} // clientID -> set(nodeId)
}

func NewInterests() *Interests {
	return &Interests{data: make(map[string]map[string]struct{})}
}

func (i *Interests) Set(clientID string, nodeIDs []string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	m := make(map[string]struct{}, len(nodeIDs))
	for _, id := range nodeIDs {
		m[id] = struct{}{}
	}
	i.data[clientID] = m
}

func (i *Interests) Get(clientID string) map[string]struct{} {
	i.mu.RLock()
	defer i.mu.RUnlock()
	src := i.data[clientID]
	if src == nil {
		return nil
	}
	dst := make(map[string]struct{}, len(src))
	for k := range src {
		dst[k] = struct{}{}
	}
	return dst
}

func (i *Interests) Delete(clientID string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	delete(i.data, clientID)
}
