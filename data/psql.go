package data

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func OpenDB() error {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s " +
		"dbname=%s sslmode=disable",
		os.Getenv("PSQL_HOST"),os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_DBNAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	DB.Close()
}
