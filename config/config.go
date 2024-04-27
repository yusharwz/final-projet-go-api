package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() (db *sql.DB, err error) {

	var PsqlInfo = "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"

	db, err = sql.Open("postgres", PsqlInfo)
	if err != nil {
		return nil, fmt.Errorf("gagal melakukan koneksi ke database %v: %v", db, err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("gagal melakukan koneksi ke database %v: %v", db, err)
	}
	return db, nil
}
