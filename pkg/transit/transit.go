package transit

type Transit interface {
	GetNamespace() string
	AfterTransporterDisconnect(err error)
	Connect(isReconnect bool) error
}
