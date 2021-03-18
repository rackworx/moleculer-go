package transporter

import (
	"fmt"

	"github.com/rackworx/moleculer-go/pkg/packet"
	"github.com/rackworx/moleculer-go/pkg/transit"
)

type Base struct {
	prefix  string
	Transit transit.Transit
}

func New(x transit.Transit) *Base {
	return &Base{
		prefix:  getPrefix(x.GetNamespace()),
		Transit: x,
	}
}

func (b *Base) GetTopicName(sub Subscription) string {
	t := fmt.Sprintf("%s.%s", b.prefix, sub.Cmd)
	if sub.NodeID != "" {
		t = fmt.Sprintf("%s.%s", t, sub.NodeID)
	}

	return t
}

func (b *Base) Receive(sub Subscription, packet packet.Packet) {
}

func getPrefix(namespace string) string {
	if namespace != "" {
		return fmt.Sprintf("MOL-%s", namespace)
	}

	return "MOL"
}
