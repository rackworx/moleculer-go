package transit

import (
	"time"

	"github.com/rackworx/moleculer-go"
	"github.com/rackworx/moleculer-go/internal/packets"
	"github.com/rackworx/moleculer-go/pkg/config"
	tx "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rs/zerolog"
)

type transit struct {
	broker      moleculer.ServiceBroker
	logger      zerolog.Logger
	transporter tx.Transporter
	connected   bool
	connecting  bool
	config      config.TransitConfig
}

func New(config config.TransitConfig, broker moleculer.ServiceBroker) *transit {
	return &transit{
		broker:      broker,
		logger:      broker.GetLogger("transit"),
		config:      config,
		transporter: config.TransporterFactory(broker.GetNamespace()),
	}
}

func (t *transit) Connect(isReconnect bool) error {
	if t.connected || t.connecting {
		return nil
	}

	t.logger.Info().Msg("connecting to transporter...")

	t.connecting = true

	for {
		err := t.transporter.Connect()

		if err != nil {
			if t.config.DisableReconnect {
				return err
			} else {
				t.logger.Warn().Err(err).Msg("connection failed")
				time.Sleep(t.config.ReconnectDelay)
				t.logger.Info().Msg("reconnecting")
				continue
			}
		}

		t.connecting = false
		t.connected = true
		t.logger.Info().Msg("connected")
		break
	}

	return nil
}

func (t *transit) getNodeID() string {
	return t.broker.GetNodeID()
}

func (t *transit) makeSubscriptions() {
	var subscriptions = []tx.Subscription{
		{
			Cmd:    packets.PACKET_EVENT,
			NodeID: t.getNodeID(),
		},
		{
			Cmd:    packets.PACKET_REQUEST,
			NodeID: t.getNodeID(),
		},
		{
			Cmd:    packets.PACKET_RESPONSE,
			NodeID: t.getNodeID(),
		},
		{
			Cmd: packets.PACKET_DISCOVER,
		},
		{
			Cmd:    packets.PACKET_DISCOVER,
			NodeID: t.getNodeID(),
		},
		{
			Cmd: packets.PACKET_INFO,
		},
		{
			Cmd:    packets.PACKET_INFO,
			NodeID: t.getNodeID(),
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
			NodeID: t.getNodeID(),
		},
		{
			Cmd:    packets.PACKET_PONG,
			NodeID: t.getNodeID(),
		},
	}

	for _, s := range subscriptions {
		t.subscribe(s)
	}
}

func (t *transit) subscribe(s tx.Subscription) {
	t.transporter.Subscribe(s)
}

func (t *transit) afterConnect(wasReconnect bool) {
	if wasReconnect {

	} else {
		t.makeSubscriptions()
	}
}
