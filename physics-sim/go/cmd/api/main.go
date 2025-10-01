package main

import (
	"log"
	"net/http"
	"physics_sim/go/pkg/api"
	"physics_sim/go/pkg/cache"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

func main() {
	cfg := loadConfig() // broker URL, topic prefix
	bus := mqttbus.New(cfg.BrokerURL, "api-gateway", cfg.LWTTopic, []byte(`{"state":"OFFLINE"}`))
	cch := cache.New()

	// Subscriptions → update caches
	bus.Subscribe(cfg.TopicPrefix+"/status/sim", func(_ mqtt.Client, m mqtt.Message) {
		var st grid_v1.SimStatus
		if err := proto.Unmarshal(m.Payload(), &st); err == nil {
			cch.UpdateSimStatus(&st)
		}
	})
	bus.Subscribe(cfg.TopicPrefix+"/node/+", func(_ mqtt.Client, m mqtt.Message) {
		var nd grid_v1.NodeDelta
		if err := proto.Unmarshal(m.Payload(), &nd); err == nil {
			cch.UpdateNode(&nd)
			// and optionally push to WS hub
		}
	})

	mux := http.NewServeMux()
	mux.Handle("/api/status", api.StatusHandler(cch))
	// add WS handler & command POSTs…

	log.Println("api listening on :8080")
	http.ListenAndServe(":8080", mux)
}
