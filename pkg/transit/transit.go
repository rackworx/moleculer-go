package transit

import "github.com/rackworx/moleculer-go/pkg/packet"

type Transit interface {
	GetNamespace() string
	AfterTransporterDisconnect(err error)
	Connect(isReconnect bool) error
	HandlePacket(packet packet.Packet)
}
