package api

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Tick     uint64 `json:"tick"`
	TopoHash string `json:"topo_hash"`
	State    string `json:"state"`
}

type Cache interface {
	SimStatus() Status
}

func StatusHandler(c Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := c.SimStatus()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)
	}
}
