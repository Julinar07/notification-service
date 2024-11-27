package mocks

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

// MockPgxPool is a mock implementation of pgxpool.Pool
type MockPgxPool struct {
	mock.Mock
}

func (m *MockPgxPool) Begin(ctx context.Context) (pgx.Tx, error) {
	args := m.Called(ctx)
	return args.Get(0).(pgx.Tx), args.Error(1)
}

func (m *MockPgxPool) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	// Capture arguments for validation
	capturedArgs := append([]interface{}{ctx, sql}, args...)
	call := m.Called(capturedArgs...)
	return call.Get(0).(pgx.Rows), call.Error(1)
}

func (m *MockPgxPool) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	// Capture arguments for validation
	capturedArgs := append([]interface{}{ctx, sql}, args...)
	call := m.Called(capturedArgs...)
	return call.Get(0).(pgx.Row)
}

func (m *MockPgxPool) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	// Capture arguments for validation
	capturedArgs := append([]interface{}{ctx, sql}, arguments...)
	call := m.Called(capturedArgs...)
	return call.Get(0).(pgconn.CommandTag), call.Error(1)
}

func (m *MockPgxPool) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	// Capture arguments for validation
	capturedArgs := append([]interface{}{ctx, tableName, columnNames, rowSrc})
	call := m.Called(capturedArgs...)
	return call.Get(0).(int64), call.Error(1)
}
