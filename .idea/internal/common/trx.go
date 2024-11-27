package common

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"sync"
)

type T struct {
	Db PgxPoolIFace
	Tx pgx.Tx
}

func RunInTx(ctx context.Context, db PgxPoolIFace, fn func(tx pgx.Tx) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	err = fn(tx)
	if err == nil {
		return tx.Commit(ctx)
	}

	rollbackErr := tx.Rollback(ctx)
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}
	return err
}

type TxFunc func() error

func RunTxInParallel(ctx context.Context, funcs ...TxFunc) []error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(funcs))

	wg.Add(len(funcs))

	for _, fn := range funcs {
		go func(fn TxFunc) {
			defer wg.Done()
			if err := fn(); err != nil {
				errCh <- err
			}
		}(fn)
	}

	wg.Wait()
	close(errCh)

	var errorList []error
	for err := range errCh {
		errorList = append(errorList, err)
	}

	return errorList
}

func runTransaction[T any](ctx context.Context, db PgxPoolIFace, fn func(ctx context.Context, conn T) error, createConn func(tx pgx.Tx) T) error {
	return RunInTx(ctx, db, func(tx pgx.Tx) error {
		conn := createConn(tx)
		return fn(ctx, conn)
	})
}
