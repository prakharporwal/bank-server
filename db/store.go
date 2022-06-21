package db

import (
	"context"
	"database/sql"
	db "github.com/prakharporwal/bank-server/db/sqlc"
	"github.com/prakharporwal/bank-server/services"
	"log"
)

//
const (
	dbSource = "postgresql://admin:password@localhost:5432/default_db?sslmode=disable"
	dbDriver = "postgres"
)

type Store struct {
	*db.Queries
	conn *sql.DB
}

var sqlInstance *Store

var klog services.Logger

func GetInstance() *Store {
	if sqlInstance == nil {
		conn, err := sql.Open(dbDriver, dbSource)
		if err != nil {
			klog.Error("connect to db failed !", err)
			panic(err)
		}
		sqlInstance = newStore(conn)
		klog.Debug("\nSuccessfully connected to database!\n")
	}
	return sqlInstance
}

func newStore(conn *sql.DB) *Store {
	return &Store{
		conn:    conn,
		Queries: db.New(conn),
	}
}

func (store *Store) Execute(statement string, args ...interface{}) error {
	_, err := store.conn.Exec(statement, args...)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (store *Store) Query(statement string, args ...interface{}) *sql.Rows {
	result, err := store.conn.Query(statement, args...)

	if err != nil {
		log.Println(err)
		return nil
	}
	// defer result.Close()

	log.Println(result.Columns())
	return result
}

func (store *Store) BeginTx(ctx context.Context, opts *sql.TxOptions) *sql.Tx {
	tx, _ := store.conn.BeginTx(ctx, opts)
	return tx
}
