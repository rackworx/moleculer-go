package config

import (
	"testing"
	"time"

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
}
