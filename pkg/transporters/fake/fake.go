package fake

import (
	"bytes"

	"github.com/olebedev/emitter"
	t "github.com/rackworx/moleculer-go/pkg/transporter"
)

var bus = &emitter.Emitter{}

type fake struct {
	*t.Base
}

func New(namespace string) t.Transporter {
	return &fake{
		Base: t.New(namespace),
	}
}

func (f *fake) Connect() error {
	return nil
}

func (f *fake) Disconnect() {
}

func (f *fake) Subscribe(subscription t.Subscription) error {
	return nil
}

func (f *fake) Send(topic string, data bytes.Buffer, meta interface{}) error {
	return nil
}
