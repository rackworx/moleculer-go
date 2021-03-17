package transporter_test

import (
	"testing"

	"github.com/rackworx/moleculer-go/internal/packets"
	tp "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rackworx/moleculer-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGetTopicNameWithoutNamespace(t *testing.T) {
	tz := &test.MockTransit{}

	tz.On("GetNamespace").Return("").Once()
	assert.Equal(t, "MOL.EVENT", tp.New(tz).GetTopicName(tp.Subscription{
		Cmd: packets.PACKET_EVENT,
	}))
}

func TestGetTopicNameWithNamespace(t *testing.T) {
	tz := &test.MockTransit{}

	tz.On("GetNamespace").Return("test").Once()
	assert.Equal(t, "MOL-test.EVENT", tp.New(tz).GetTopicName(tp.Subscription{
		Cmd: packets.PACKET_EVENT,
	}))
}

func TestGetTopicNameWithNodeID(t *testing.T) {
	tz := &test.MockTransit{}

	tz.On("GetNamespace").Return("").Once()
	assert.Equal(t, "MOL.EVENT.test", tp.New(tz).GetTopicName(tp.Subscription{
		Cmd:    packets.PACKET_EVENT,
		NodeID: "test",
	}))

}
