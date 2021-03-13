package moleculer

import (
	"bytes"

	"github.com/rs/zerolog"
)

type ServiceBroker interface {
	Start()
	GetLogger(name string) zerolog.Logger
	GetNamespace() string
	GetNodeID() string
}

type Transit interface {
	GetBroker() ServiceBroker
}

type Subscription struct {
	// The command to subscribe to
	Cmd string
	// The node to subscribe to
	NodeID string
}

type Transporter interface {
	// to the transporter bus
	Connect() error
	// from the transporter bus
	Disconnect()
	// to a command
	Subscribe(Subscription) error
	// data buffer
	Send(topic string, data bytes.Buffer, meta interface{}) error
	// starts the transporter and waits for an error

}
