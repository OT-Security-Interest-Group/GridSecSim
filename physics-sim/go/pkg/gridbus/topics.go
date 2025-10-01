package mqttbus

import "path"

// Topic helpers keep strings centralized. Prefix like "grid/v1".
type Topics struct{ Prefix string }

func (t Topics) StatusSim() string               { return path.Join(t.Prefix, "status/sim") }
func (t Topics) StatusComp(compID string) string { return path.Join(t.Prefix, "status/comp", compID) }

func (t Topics) Topology() string { return path.Join(t.Prefix, "topology") }

func (t Topics) Announce() string            { return path.Join(t.Prefix, "announce") }
func (t Topics) Ack(compID string) string    { return path.Join(t.Prefix, "ack", compID) }
func (t Topics) Intent(compID string) string { return path.Join(t.Prefix, "intent", compID) }
func (t Topics) IntentConfirm(compID string) string {
	return path.Join(t.Prefix, "intent/confirm", compID)
}

func (t Topics) Measure(compID string) string { return path.Join(t.Prefix, "measure", compID) }
func (t Topics) Cmd(compID string) string     { return path.Join(t.Prefix, "cmd", compID) }
func (t Topics) CmdAck(compID string) string  { return path.Join(t.Prefix, "cmdack", compID) }

func (t Topics) Node(nodeID string) string { return path.Join(t.Prefix, "node", nodeID) }
func (t Topics) NodeAll() string           { return path.Join(t.Prefix, "node", "+") }
func (t Topics) Edge(edgeID string) string { return path.Join(t.Prefix, "edge", edgeID) }
func (t Topics) EdgeAll() string           { return path.Join(t.Prefix, "edge", "+") }
