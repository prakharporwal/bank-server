package db

import (
	"database/sql"
	"fmt"
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

func (store *Store) Query(statement string, args ...interface{}) {
	result, err := store.db.Exec(statement, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
