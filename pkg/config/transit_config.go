package config

import (
	"time"

	"github.com/imdario/mergo"
	t "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rackworx/moleculer-go/pkg/transporters/fake"
)

type TransitConfig struct {
	MaxQueueSize        int
	MaxChunkSize        int
	DisableReconnect    bool
	DisableVersionCheck bool
	ReconnectDelay      time.Duration
	TransporterFactory  t.TransporterFactory
}

func createDefaultTransitConfig(config TransitConfig) TransitConfig {
	cfg := TransitConfig{
		MaxQueueSize:       50000,
		MaxChunkSize:       256000,
		ReconnectDelay:     5 * time.Second,
		TransporterFactory: fake.New(),
	}

	mergo.Merge(&config, cfg)

	return config
}
