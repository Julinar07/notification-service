package mocks

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

type MockRows struct {
	mock.Mock
}

func (m *MockRows) Next() bool {
	args := m.MethodCalled("Next")
	return args.Bool(0)
}

func (m *MockRows) Close() {
	m.Called()
}

func (m *MockRows) Err() error {
	args := m.MethodCalled("Err")
	return args.Error(0)
}

func (m *MockRows) Scan(dest ...any) error {
	args := m.Called(dest...)
	return args.Error(0)
}

func (m *MockRows) CommandTag() pgconn.CommandTag {
	m.Called()
	return pgconn.CommandTag{}
}

func (m *MockRows) CommandName() string {
	return ""
}

func (m *MockRows) CommandArgs() []interface{} {
	m.Called()
	return nil
}

func (m *MockRows) Columns() []string {
	m.Called()
	return nil
}

func (m *MockRows) RawValues() [][]byte {
	m.Called()
	return nil
}

func (m *MockRows) Values() ([]interface{}, error) {
	m.Called()
	return nil, nil
}

func (m *MockRows) Conn() *pgx.Conn {
	m.Called()
	return nil
}

func (m *MockRows) FieldDescriptions() []pgconn.FieldDescription {
	m.Called()
	return nil
}
