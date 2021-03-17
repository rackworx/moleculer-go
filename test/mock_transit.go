package test

import "github.com/stretchr/testify/mock"

type MockTransit struct {
	mock.Mock
}

func (m *MockTransit) GetNamespace() string {
	args := m.Called()

	return args.String(0)
}

func (m *MockTransit) AfterTransporterDisconnect(err error) {
	m.Called(err)
}

func (m *MockTransit) Connect(isReconnect bool) error {
	args := m.Called(isReconnect)

	return args.Error(0)
}
