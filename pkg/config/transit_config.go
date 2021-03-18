package config

import (
	"time"

	"github.com/imdario/mergo"
	"github.com/rackworx/moleculer-go/pkg/serializer"
	"github.com/rackworx/moleculer-go/pkg/serializers/json"
	t "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rackworx/moleculer-go/pkg/transporters/fake"
)

type TransitConfig struct {
	MaxQueueSize        int
	MaxChunkSize        int
	DisableReconnect    bool
	DisableVersionCheck bool
	ReconnectDelay      time.Duration
	ConnectTimeout      time.Duration
	TransporterFactory  t.TransporterFactory
	Serializer          serializer.Serializer
}

func createDefaultTransitConfig(config TransitConfig) TransitConfig {
	cfg := TransitConfig{
		MaxQueueSize:       50000,
		MaxChunkSize:       256000,
		ReconnectDelay:     5 * time.Second,
		TransporterFactory: fake.NewFactory(),
		Serializer:         json.New(),
		ConnectTimeout:     10 * time.Second,
	}

	mergo.Merge(&config, cfg)

	return config
}
