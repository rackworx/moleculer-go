package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rackworx/moleculer-go/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigDefaults(t *testing.T) {
	cfg := config.New(config.Config{})

	hostname, _ := os.Hostname()

	assert.Equal(t, "", cfg.Namespace)
	assert.Equal(t, fmt.Sprintf("%s%d", hostname, os.Getpid()), cfg.NodeID)

	assert.NotEqual(t, config.LogConfig{}, cfg.Logging)
	assert.NotEqual(t, config.TransitConfig{}, cfg.Transit)
}
