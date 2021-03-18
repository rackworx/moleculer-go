package config

import (
	"reflect"
	"testing"
	"time"

	"github.com/rackworx/moleculer-go/pkg/serializers/json"
	"github.com/stretchr/testify/assert"
)

func TestCreateDefaultTransitConfig(t *testing.T) {
	cfg := createDefaultTransitConfig(TransitConfig{})
	assert.Equal(
		t,
		50000,
		cfg.MaxQueueSize,
	)
	assert.Equal(
		t,
		256000,
		cfg.MaxChunkSize,
	)

	assert.Equal(
		t,
		5*time.Second,
		cfg.ReconnectDelay,
	)
	assert.Equal(
		t,
		500,
		createDefaultTransitConfig(TransitConfig{MaxQueueSize: 500}).MaxQueueSize,
	)
	assert.Equal(
		t,
		reflect.TypeOf(json.New()),
		reflect.TypeOf(createDefaultTransitConfig(TransitConfig{}).Serializer),
	)
	assert.Equal(
		t,
		1*time.Millisecond,
		createDefaultTransitConfig(
			TransitConfig{ReconnectDelay: 1 * time.Millisecond},
		).ReconnectDelay,
	)
}
