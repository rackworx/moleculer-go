package transporter_test

import (
	"testing"

	tp "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rackworx/moleculer-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGetTopicName(t *testing.T) {
	tz := &test.MockTransit{}

	tz.On("GetNamespace").Return("").Once()
	assert.Equal(t, "MOL.test", tp.New(tz).GetTopicName("test"))

	tz.On("GetNamespace").Return("test").Once()
	assert.Equal(t, "MOL-test.test", tp.New(tz).GetTopicName("test"))
}

func TestGetTopicNameForNode(t *testing.T) {
	tz := &test.MockTransit{}

	tz.On("GetNamespace").Return("").Once()

	assert.Equal(
		t,
		"MOL.test.node",
		tp.New(tz).GetTopicNameForNode("test", "node"),
	)
}
