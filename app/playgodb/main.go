package main

import (
	"database/sql"
	"errors"

	"github.com/dokyan1989/g1/app/playgodb/store"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := initializeDB()
	if err != nil {
		panic(err.Error())
	}

	var s store.Store
	s = store.NewStore(db)
	defer s.Close()
	if err = s.Ping(); err != nil {
		panic("unable to connect to database")
	}
}

func initializeDB() (*sql.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New("unable to use data source name")
	}

	return db, nil
}
