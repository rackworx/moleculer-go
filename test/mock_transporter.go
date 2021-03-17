package test

import (
	"bytes"

	"github.com/rackworx/moleculer-go/pkg/transporter"
	"github.com/stretchr/testify/mock"
)

type MockTransporter struct {
	mock.Mock
}

func (m *MockTransporter) Connect() error {
	args := m.Called()

	return args.Error(0)
}

func (m *MockTransporter) Disconnect() {}

func (m *MockTransporter) Subscribe(sub transporter.Subscription) error {
	args := m.Called(sub)

	return args.Error(0)
}

func (m *MockTransporter) Send(topic string, data bytes.Buffer, meta interface{}) error {
	return nil
}

func (m *MockTransporter) IsConnected() bool {
	args := m.Called()

	return args.Bool(0)
}
