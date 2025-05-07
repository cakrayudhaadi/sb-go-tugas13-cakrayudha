package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "sanbercode_golang"
)

var DBConnection *sql.DB

func Initiator() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	DBConnection = db
	err = DBConnection.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
