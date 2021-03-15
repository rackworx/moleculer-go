package test

import (
	"github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/mock"
)

type MockBroker struct {
	mock.Mock
}

func (m *MockBroker) Start() {}
func (m *MockBroker) GetLogger(_ string) zerolog.Logger {
	return log.With().Logger().Level(zerolog.Disabled)
}

func (m *MockBroker) GetNamespace() string {
	args := m.Called()

	return args.String(0)
}

func (m *MockBroker) GetNodeID() string {
	return ""
}

func (m *MockBroker) GetTransporter() transporter.Transporter {
	return nil
}
