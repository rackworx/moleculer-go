package transporter

import "bytes"

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
