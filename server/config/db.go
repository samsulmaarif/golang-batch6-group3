package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	DB_HOST = "127.0.0.1"
	DB_PORT = "5432"
	DB_USER = "nest"
	DB_PASS = "rahasia"
	DB_NAME = "golangnest"
)

func ConnectDB() *sql.DB {

	host := DB_HOST
	port := DB_PORT
	user := DB_USER
	password := DB_PASS
	dbname := DB_NAME

	dbs, err := getPostgres(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}

	err = dbs.Ping()
	if err != nil {
		panic(err)
	}

	return dbs

}

func getPostgres(host, port, user, password, dbname string) (*sql.DB, error) {
	desc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbs, err := connectDB(desc)
	if err != nil {

		return nil, err
	}

	return dbs, nil

}

func connectDB(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
