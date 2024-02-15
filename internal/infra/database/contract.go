package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lccmrx/rinha-bank/internal/repository"
)

type DataManager interface {
	Repo
	Begin() (TransactionManager, error)
	Close() error
}

// TransactionManager holds the methods that manipulates the main
// data, from within a transaction.
type TransactionManager interface {
	Repo
	Rollback() error
	Commit() error
	GetDBTransaction() *sqlx.Tx
}

type Repo interface {
	Client() repository.Client
	Transaction() repository.Transaction
}

type Executor interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
	Rebind(string) string
}
