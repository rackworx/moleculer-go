package transit_test

import (
	"errors"
	"testing"
	"time"

	"github.com/rackworx/moleculer-go/internal/packets"
	"github.com/rackworx/moleculer-go/internal/transit"
	"github.com/rackworx/moleculer-go/pkg/config"
	tz "github.com/rackworx/moleculer-go/pkg/transit"
	tx "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rackworx/moleculer-go/test"
	"github.com/stretchr/testify/assert"
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

func TestInitialConnect(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(_ tz.Transit) tx.Transporter {
		return transporter
	}

	broker.On("GetNodeID").Return("test")

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			TransporterFactory: transporterFactory,
		},
	})

	subscriptions := []tx.Subscription{
		{
			Cmd:    packets.PACKET_EVENT,
			NodeID: "test",
		},
		{
			Cmd:    packets.PACKET_REQUEST,
			NodeID: "test",
		},
		{
			Cmd:    packets.PACKET_RESPONSE,
			NodeID: "test",
		},
		{
			Cmd: packets.PACKET_DISCOVER,
		},
		{
			Cmd:    packets.PACKET_DISCOVER,
			NodeID: "test",
		},
		{
			Cmd: packets.PACKET_INFO,
		},
		{
			Cmd:    packets.PACKET_INFO,
			NodeID: "test",
		},
		{
			Cmd: packets.PACKET_DISCONNECT,
		},
		{
			Cmd: packets.PACKET_HEARTBEAT,
		},
		{
			Cmd: packets.PACKET_PING,
		},
		{
			Cmd:    packets.PACKET_PING,
			NodeID: "test",
		},
		{
			Cmd:    packets.PACKET_PONG,
			NodeID: "test",
		},
	}

	for _, v := range subscriptions {
		transporter.On("Subscribe", v).Return(nil).Once()
	}

	transporter.On("Connect").Return(nil)
	transporter.On("IsConnected").Return(true).Once()

	transit := transit.New(cfg.Transit, broker)
	transit.Connect(false)

	broker.AssertExpectations(t)
	transporter.AssertExpectations(t)
}

func TestTransporterSubscribeReturnsError(t *testing.T) {
	broker := &test.MockBroker{}
	transporter := &test.MockTransporter{}
	transporterFactory := func(_ tz.Transit) tx.Transporter {
		return transporter
	}

	cfg := config.New(config.Config{
		Transit: config.TransitConfig{
			TransporterFactory: transporterFactory,
		},
	})

	broker.On("GetNodeID").Return("test")

	transporter.On("Connect").Return(nil)
	transporter.On("IsConnected").Return(true)

	transporter.On("Subscribe", tx.Subscription{
		Cmd:    packets.PACKET_EVENT,
		NodeID: "test",
	}).Return(assert.AnError)

	transit := transit.New(cfg.Transit, broker)

	assert.Error(t, assert.AnError, transit.Connect(false))
}

func TestGetNamespace(t *testing.T) {
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

	broker.On("GetNamespace").Return("test")

	transit := transit.New(cfg.Transit, broker)

	assert.Equal(t, "test", transit.GetNamespace())
}
