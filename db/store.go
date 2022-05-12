package db

import (
	"context"
	"database/sql"
	"log"
)

type Store struct {
	// *Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) Execute(statement string, args ...interface{}) sql.Result {
	result, err := store.db.Exec(statement, args...)

	if err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func (store *Store) Query(statement string, args ...interface{}) *sql.Rows {
	result, err := store.db.Query(statement, args...)

	if err != nil {
		log.Println(err)
		return nil
	}
	// defer result.Close()

	log.Println(result.Columns())

	return result
}

func (store *Store) BeginTx(ctx context.Context, opts *sql.TxOptions) *sql.Tx {
	tx, _ := store.db.BeginTx(ctx, opts)
	return tx
}
