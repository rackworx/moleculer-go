package transporter

import (
	"fmt"
)

type Base struct {
	prefix string
}

func New(namespace string) *Base {
	return &Base{
		prefix: getPrefix(namespace),
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
