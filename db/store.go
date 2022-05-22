package db

import (
	"database/sql"
	"fmt"
)

type Store interface {
	// *Queries
	Query(statement string, arg ...interface{})
	GetQuery(statement string, arg ...interface{})
}

type SQLStore struct {
	Store
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (sqlStore *SQLStore) Query(statement string, args ...interface{}) {
	result, err := sqlStore.db.Exec(statement, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func (sqlStore *SQLStore) GetQuery(statement string, args ...interface{}) error {
	sqlStore.Query(statement, args)
	return nil
}
