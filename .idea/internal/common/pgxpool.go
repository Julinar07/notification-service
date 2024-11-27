package common

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPoolIFace interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type PgxPoolWrapper struct {
	Pool *pgxpool.Pool
}

func (w *PgxPoolWrapper) Begin(ctx context.Context) (pgx.Tx, error) {
	return w.Pool.Begin(ctx)
}

func (w *PgxPoolWrapper) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return w.Pool.Query(ctx, sql, args...)
}

func (w *PgxPoolWrapper) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return w.Pool.QueryRow(ctx, sql, args...)
}

func (w *PgxPoolWrapper) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return w.Pool.Exec(ctx, sql, arguments...)
}

func (w *PgxPoolWrapper) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return w.Pool.CopyFrom(ctx, tableName, columnNames, rowSrc)
}
