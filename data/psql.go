package data

import (
	"database/sql"

	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func OpenDB(host, port, user, password, name string) error {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s " +
		"dbname=%s sslmode=disable", host, port, user, password, name)

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
