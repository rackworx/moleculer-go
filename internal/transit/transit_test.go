package transit_test

import (
	"errors"
	"testing"
	"time"

	"github.com/rackworx/moleculer-go/internal/transit"
	"github.com/rackworx/moleculer-go/pkg/config"
	tz "github.com/rackworx/moleculer-go/pkg/transit"
	tx "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rackworx/moleculer-go/test"
)

func TestTransitCleanConnect(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(tz tz.Transit) tx.Transporter {
		return transporter
	}

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			TransporterFactory: transporterFactory,
		},
	})

	broker.On("GetNamespace").Return("")
	transporter.On("Connect").Return(nil)

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(false)

	transporter.AssertExpectations(t)
}

func TestTransitInitialReconnect(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(_ tz.Transit) tx.Transporter {
		return transporter
	}

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			ReconnectDelay:     1 * time.Millisecond,
			TransporterFactory: transporterFactory,
		},
	})

	err := errors.New("test")

	broker.On("GetNamespace").Return("")

	transporter.On("Connect").Return(err).Once()
	transporter.On("Connect").Return(nil).Once()

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(false)

	transporter.AssertExpectations(t)
}

func TestTransitInitialNoReconnect(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(_ tz.Transit) tx.Transporter {
		return transporter
	}

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			DisableReconnect:   true,
			TransporterFactory: transporterFactory,
		},
	})

	err := errors.New("test")

	broker.On("GetNamespace").Return("").Once()
	transporter.On("Connect").Return(err).Once()

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(false)

	transporter.AssertExpectations(t)
}
