package db

import (
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://admin:password@localhost:5432/default_db?sslmode=disable"
	dbDriver = "postgres"
)

var testQueries *Queries

//func TestMain(m *testing.M) {
//	conn, err := sql.Open(dbDriver, dbSource)
//	if err != nil {
//		klog.Info("connect to db failed !", err)
//		panic(err)
//	}
//	defer conn.Close()
//
//	klog.Debug("\nSuccessfully connected to database!\n")
//
//	//store := db.NewStore(conn)
//	testQueries = New(conn)
//
//	os.Exit(m.Run())
//}
