package postgres

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lccmrx/rinha-bank/internal/infra/config"
	"github.com/lccmrx/rinha-bank/internal/infra/database"
	"github.com/lccmrx/rinha-bank/internal/infra/database/postgres/model"
	"github.com/lccmrx/rinha-bank/internal/repository"
	_ "github.com/lib/pq"
)

var (
	instance     *PostgresConnector
	dbInstance   *sqlx.DB
	onceDB       sync.Once
	onceInstance sync.Once
	connErr      error
)

type PostgresConnector struct {
	db *sqlx.DB

	client      repository.Client
	transaction repository.Transaction
}

var _ database.DataManager = (*PostgresConnector)(nil)

func Instance(ctx context.Context, cfg *config.Config) (*PostgresConnector, error) {
	onceInstance.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Database)
		db, err := GetDB(ctx, dsn)
		if err != nil {
			connErr = err
			return
		}

		instance = &PostgresConnector{db: db}

		instance.client = &model.Client{Conn: db}
		instance.transaction = &model.Transaction{Conn: db}
	})

	return instance, connErr
}

func GetDB(ctx context.Context, dsn string) (*sqlx.DB, error) {
	onceDB.Do(func() {
		connection, err := sqlx.ConnectContext(ctx, "postgres", dsn)
		if err != nil {
			panic(fmt.Errorf("couldn't start database: %s", err))
		}

		connection.SetConnMaxLifetime(time.Minute)
		connection.SetMaxIdleConns(25)
		connection.SetMaxOpenConns(25)

		err = connection.Ping()
		if err != nil {
			connErr = err
			return
		}

		dbInstance = connection
	})

	return dbInstance, connErr
}

func (c *PostgresConnector) Begin() (database.TransactionManager, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		return nil, err
	}

	return newTransaction(tx), nil
}

func (c *PostgresConnector) Close() (err error) {
	return c.db.Close()
}

func (c *PostgresConnector) Client() repository.Client {
	return c.client
}

func (c *PostgresConnector) Transaction() repository.Transaction {
	return c.transaction
}

type transaction struct {
	tx         *sqlx.Tx
	committed  bool
	rolledback bool

	client      *model.Client
	transaction *model.Transaction
}

func newTransaction(tx *sqlx.Tx) *transaction {
	instance := &transaction{tx: tx}

	instance.client = &model.Client{Conn: tx}
	instance.transaction = &model.Transaction{Conn: tx}

	return instance
}

// Commit persists changes to database
func (t *transaction) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}

	t.committed = true

	return nil
}

// Rollback discards changes on database
func (t *transaction) Rollback() error {
	if t != nil && !t.committed && !t.rolledback {
		err := t.tx.Rollback()
		if err != nil {
			return err
		}

		t.rolledback = true
	}

	return nil
}

// GetDBTransaction returns the transaction instance reference
func (t *transaction) GetDBTransaction() *sqlx.Tx {
	return t.tx
}

// States returns the states set
func (t *transaction) Client() repository.Client {
	return t.client
}

// Cities returns the cities set
func (t *transaction) Transaction() repository.Transaction {
	return t.transaction
}
