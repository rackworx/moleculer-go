package base_test

import (
	"testing"

	"github.com/rackworx/moleculer-go/pkg/transporters/base"
	"github.com/stretchr/testify/assert"
)

func TestGetTopicName(t *testing.T) {
	assert.Equal(t, "MOL.test", base.New("").GetTopicName("test"))
	assert.Equal(t, "MOL-test.test", base.New("test").GetTopicName("test"))
}

func TestGetTopicNameForNode(t *testing.T) {
	assert.Equal(
		t,
		"MOL.test.node",
		base.New("").GetTopicNameForNode("test", "node"),
	)
}
