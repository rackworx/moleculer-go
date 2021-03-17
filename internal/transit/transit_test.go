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

func TestTransitConnectReconnectWithDisabledReconnect(t *testing.T) {
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

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(true)

	transporter.AssertExpectations(t)
}

func TestTransitConnectReconnectWithoutDisabledReconnect(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(_ tz.Transit) tx.Transporter {
		return transporter
	}

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			TransporterFactory: transporterFactory,
			ReconnectDelay:     1 * time.Millisecond,
		},
	})

	transporter.On("Connect").Return(nil)
	transporter.On("IsConnected").Return(true).Once()

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(true)

	transporter.AssertExpectations(t)
}

func TestTransporterConnectError(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(_ tz.Transit) tx.Transporter {
		return transporter
	}

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			TransporterFactory: transporterFactory,
			ReconnectDelay:     1 * time.Millisecond,
		},
	})

	transporter.On("Connect").Return(errors.New("error!")).Once()
	transporter.On("Connect").Return(nil)
	transporter.On("IsConnected").Return(true).Once()

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(true)

	transporter.AssertExpectations(t)
}

func TestConnectTimeout(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(_ tz.Transit) tx.Transporter {
		return transporter
	}

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			TransporterFactory: transporterFactory,
			ReconnectDelay:     1 * time.Millisecond,
			ConnectTimeout:     10 * time.Millisecond,
		},
	})

	transporter.On("Connect").Return(nil)
	transporter.On("IsConnected").Return(false).Times(11)
	transporter.On("IsConnected").Return(true).Once()

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(true)

	transporter.AssertExpectations(t)
}
