package fake

import (
	"bytes"

	"github.com/olebedev/emitter"
	"github.com/rackworx/moleculer-go"
	"github.com/rackworx/moleculer-go/pkg/transporters/base"
)

var bus = &emitter.Emitter{}

type fake struct {
	*base.Base
}

func New(namespace string) moleculer.Transporter {
	return &fake{
		Base: base.New(namespace),
	}
}

func (f *fake) Connect() error {
	return nil
}

func (f *fake) Disconnect() {
}

func (f *fake) Subscribe(subscription moleculer.Subscription) error {
	return nil
}

func (f *fake) Send(topic string, data bytes.Buffer, meta interface{}) error {
	return nil
}
