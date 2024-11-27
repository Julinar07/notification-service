package mocks

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

// MockTx is a mock implementation of pgx.Tx
type MockTx struct {
	mock.Mock
}

func (m *MockTx) Begin(ctx context.Context) (pgx.Tx, error) {
	args := m.Called(ctx)
	return args.Get(0).(pgx.Tx), args.Error(1)
}

func (m *MockTx) Commit(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockTx) Rollback(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockTx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	args := m.Called(ctx, tableName, columnNames, rowSrc)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	args := m.Called(ctx, b)
	return args.Get(0).(pgx.BatchResults)
}

func (m *MockTx) LargeObjects() pgx.LargeObjects {
	args := m.Called()
	return args.Get(0).(pgx.LargeObjects)
}

func (m *MockTx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	args := m.Called(ctx, name, sql)
	return args.Get(0).(*pgconn.StatementDescription), args.Error(1)
}

func (m *MockTx) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	args := m.Called(ctx, sql, arguments)
	return args.Get(0).(pgconn.CommandTag), args.Error(1)
}

func (m *MockTx) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	args = append([]any{ctx, sql}, args...)
	call := m.Called(args...)
	return call.Get(0).(pgx.Rows), call.Error(1)
}

func (m *MockTx) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	args = append([]any{ctx, sql}, args...)
	call := m.Called(args...)
	return call.Get(0).(pgx.Row)
}

func (m *MockTx) Conn() *pgx.Conn {
	args := m.Called()
	return args.Get(0).(*pgx.Conn)
}

// MockRunInTx is a mock implementation of the RunInTx function
type MockRunInTx struct {
	mock.Mock
}

func (m *MockRunInTx) RunInTx(ctx context.Context, db *MockPgxPool, fn func(tx pgx.Tx) error) error {
	args := m.Called(ctx, db, fn)

	//// Execute the provided function
	//tx, err := db.Begin(ctx)
	//if err != nil {
	//	return err
	//}
	//
	//// Simulate function execution (using the actual fn passed or from mock)
	//executionErr := fn(tx)
	//if executionErr == nil {
	//	commitErr := tx.Commit(ctx)
	//	if commitErr != nil {
	//		return commitErr
	//	}
	//} else {
	//	rollbackErr := tx.Rollback(ctx)
	//	if rollbackErr != nil {
	//		return errors.Join(executionErr, rollbackErr)
	//	}
	//	return executionErr
	//}

	return args.Error(0)
}
