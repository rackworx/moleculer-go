package fake_test

import (
	"testing"

	"github.com/rackworx/moleculer-go/pkg/config"
	"github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rackworx/moleculer-go/pkg/transporters/fake"
	"github.com/rackworx/moleculer-go/test"
	"github.com/stretchr/testify/assert"
)

var transit *test.MockTransit

func createFakeTransporter() transporter.Transporter {
	transit = &test.MockTransit{}

	transit.On("GetNamespace").Return("")
	transit.On("AfterTransporterDisconnect", nil).Return()

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			TransporterFactory: fake.NewFactory(),
		},
	})

	return cfg.Transit.TransporterFactory(transit)
}

func TestConnect(t *testing.T) {
	f := createFakeTransporter()

	assert.Nil(t, f.Connect())
}

func TestDisconnect(t *testing.T) {
	f := createFakeTransporter()

	f.Disconnect()

	transit.AssertExpectations(t)
}

func TestIsConnected(t *testing.T) {
	f := createFakeTransporter()

	assert.False(t, f.IsConnected())

	f.Connect()

	assert.True(t, f.IsConnected())
}
