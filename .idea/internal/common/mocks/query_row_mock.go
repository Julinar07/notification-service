package mocks

import "github.com/stretchr/testify/mock"

type MockRow struct {
	mock.Mock
}

func (m *MockRow) Scan(dest ...any) error {
	args := m.Called(dest...)
	return args.Error(0)
}

func (m *MockRow) Close() {
	m.Called()
}
