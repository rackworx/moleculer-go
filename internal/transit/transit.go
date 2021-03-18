package transit

import (
	"errors"
	"time"

	"github.com/rackworx/moleculer-go"
	"github.com/rackworx/moleculer-go/internal/packets"
	"github.com/rackworx/moleculer-go/pkg/config"
	"github.com/rackworx/moleculer-go/pkg/packet"
	tz "github.com/rackworx/moleculer-go/pkg/transit"
	tx "github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rs/zerolog"
)

type transit struct {
	broker      moleculer.ServiceBroker
	logger      zerolog.Logger
	transporter tx.Transporter
	config      config.TransitConfig
}

func New(config config.TransitConfig, broker moleculer.ServiceBroker) tz.Transit {
	xit := &transit{
		broker: broker,
		logger: broker.GetLogger("transit"),
		config: config,
	}

	xit.transporter = config.TransporterFactory(xit)

	return xit
}

func (t *transit) Connect(isReconnect bool) error {
	t.logger.Info().Msg("connecting to transporter")

	if isReconnect && t.config.DisableReconnect {
		return errors.New("transported failed to connect and DisableReconnect " +
			"is set to true")
	}

	if isReconnect {
		time.Sleep(t.config.ReconnectDelay)
	}

	err := t.transporter.Connect()

	if err != nil {
		t.AfterTransporterDisconnect(err)
		return nil
	}

	timeout := 0 * time.Millisecond
	for ok := true; ok; ok = !t.transporter.IsConnected() {
		if timeout > t.config.ConnectTimeout {
			t.AfterTransporterDisconnect(
				errors.New("timed out connecting to transporter"),
			)
			return nil
		}

		time.Sleep(10 * time.Millisecond)
		timeout = timeout + 10*time.Millisecond
	}

	err = t.afterTransporterConnect(isReconnect)

	if err != nil {
		return err
	}

	return nil
}

func (t *transit) HandlePacket(packet packet.Packet) {
	data := packet.Payload.([]byte)

	switch packet.Type {
	case packets.PACKET_INFO:
		payload := &packets.InfoPayload{}

		err := t.config.Serializer.Unmarshal(data, payload)

		if err != nil {
			t.logger.Error().Err(err).Msg("")
		}
	}
}

func (t *transit) AfterTransporterDisconnect(err error) {
	t.logger.Error().Err(err).Msg("")
	t.Connect(true)
}

func (t *transit) afterTransporterConnect(reconnect bool) error {
	if reconnect {
		// send local node info
		return nil
	} else {
		err := t.makeSubscriptions()

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *transit) getNodeID() string {
	return t.broker.GetNodeID()
}

func (t *transit) makeSubscriptions() error {
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
		err := t.subscribe(s)

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *transit) GetNamespace() string {
	return t.broker.GetNamespace()
}

func (t *transit) subscribe(s tx.Subscription) error {
	err := t.transporter.Subscribe(s)

	if err != nil {
		return err
	}

	return nil
}
