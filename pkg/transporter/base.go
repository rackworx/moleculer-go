package transporter

import (
	"fmt"

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

func (b *Base) GetTopicNameForNode(cmd string, nodeID string) string {
	return b.GetTopicName(fmt.Sprintf("%s.%s", cmd, nodeID))
}

func (b *Base) GetTopicName(cmd string) string {
	return fmt.Sprintf("%s.%s", b.prefix, cmd)
}

func getPrefix(namespace string) string {
	if namespace != "" {
		return fmt.Sprintf("MOL-%s", namespace)
	}

	return "MOL"
}
