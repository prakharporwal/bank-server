package db

import (
	"context"
	"database/sql"
	"log"
)

type Store struct {
	// *Queries
	conn *sql.DB
}

func NewStore(conn *sql.DB) *Store {
	return &Store{
		conn: conn,
	}
}

func (store *Store) Execute(statement string, args ...interface{}) sql.Result {
	result, err := store.conn.Exec(statement, args...)

	if err != nil {
		log.Println(err)
		return nil
	}

	return result
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
