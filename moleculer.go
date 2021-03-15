package moleculer

import (
	"github.com/rs/zerolog"
)

type ServiceBroker interface {
	Start()
	GetLogger(name string) zerolog.Logger
	GetNamespace() string
	GetNodeID() string
}

type Transit interface {
	GetBroker() ServiceBroker
}


