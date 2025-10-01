package gridbus

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Bus struct {
	cli mqtt.Client
}

func New(url, clientID string, lwtTopic string, lwtPayload []byte) *Bus {
	opts := mqtt.NewClientOptions().AddBroker(url).SetClientID(clientID)
	opts.SetWill(lwtTopic, string(lwtPayload), 1, true)
	opts.SetOrderMatters(false)
	opts.SetAutoReconnect(true)
	opts.SetConnectTimeout(5 * time.Second)
	cli := mqtt.NewClient(opts)
	tok := cli.Connect()
	tok.Wait()
	if tok.Error() != nil {
		log.Fatalf("mqtt connect: %v", tok.Error())
	}
	return &Bus{cli: cli}
}

func (b *Bus) Subscribe(topic string, cb mqtt.MessageHandler) {
	tok := b.cli.Subscribe(topic, 1, cb)
	tok.Wait()
	if tok.Error() != nil {
		log.Fatalf("sub %s: %v", topic, tok.Error())
	}
}

func (b *Bus) Publish(topic string, payload []byte, retain bool) error {
	tok := b.cli.Publish(topic, 1, retain, payload)
	tok.Wait()
	return tok.Error()
}
