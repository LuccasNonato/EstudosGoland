package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "123456"
	DBName   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBName)

	db, err := sql.Open("postgres", psqlinfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!" + DBName)
	return db, nil
}
