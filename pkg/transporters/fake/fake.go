package fake

import (
	"bytes"

	"github.com/olebedev/emitter"
	"github.com/rackworx/moleculer-go/pkg/transit"
	t "github.com/rackworx/moleculer-go/pkg/transporter"
)

var bus = &emitter.Emitter{}

type fake struct {
	*t.Base
	connected bool
}

func New() t.TransporterFactory {
	return func(x transit.Transit) t.Transporter {
		return &fake{
			Base: t.New(x),
		}
	}
}

func (f *fake) Connect() error {
	f.connected = true

	return nil
}

func (f *fake) Disconnect() {
	f.Base.Transit.AfterTransporterDisconnect(nil)
}

func (f *fake) Subscribe(subscription t.Subscription) error {
	return nil
}

func (f *fake) Send(topic string, data bytes.Buffer, meta interface{}) error {
	return nil
}

func (f *fake) IsConnected() bool {
	return f.connected
}
