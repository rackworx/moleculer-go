package transit

type Transit interface {
	GetNamespace() string
	AfterTransporterConnect(reconnect bool)
	AfterTransporterDisconnect(err error)
	Connect(isReconnect bool)
}
