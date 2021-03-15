package transporter_test

import (
	"testing"

	tp "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/stretchr/testify/assert"
)

func TestGetTopicName(t *testing.T) {
	assert.Equal(t, "MOL.test", tp.New("").GetTopicName("test"))
	assert.Equal(t, "MOL-test.test", tp.New("test").GetTopicName("test"))
}

func TestGetTopicNameForNode(t *testing.T) {
	assert.Equal(
		t,
		"MOL.test.node",
		tp.New("").GetTopicNameForNode("test", "node"),
	)
}
