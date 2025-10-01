package control

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"

	"yourmod/go/pkg/mqttbus"
	grid "yourmod/go/pkg/pb/grid/v1"
)

// Publish a Command protobuf to a component's command topic.
func SendSetOnOff(cli mqtt.Client, topics mqttbus.Topics, compID string, on bool, applyTick uint64) error {
	cmd := &grid.Command{
		Env:       &grid.Envelope{}, // fill env if you want seq/msg_id here
		CompId:    compID,
		Type:      "SET_ON",
		Args:      map[string]string{"value": fmt.Sprintf("%t", on)},
		ApplyTick: applyTick,
	}
	payload, err := proto.Marshal(cmd)
	if err != nil {
		return err
	}
	token := cli.Publish(topics.Cmd(compID), 1, false, payload)
	token.Wait()
	return token.Error()
}

// Generic helper for arbitrary commands
func SendCommand(cli mqtt.Client, topics mqttbus.Topics, cmd *grid.Command) error {
	b, err := proto.Marshal(cmd)
	if err != nil {
		return err
	}
	tok := cli.Publish(topics.Cmd(cmd.GetCompId()), 1, false, b)
	tok.Wait()
	return tok.Error()
}
