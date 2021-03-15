package transporter

type Subscription struct {
	// The command to subscribe to
	Cmd string
	// The node to subscribe to
	NodeID string
}
