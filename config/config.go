package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() (db *sql.DB, err error) {

	db, err = sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return db, nil
}
