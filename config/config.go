package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() (db *sql.DB, err error) {

	var PsqlInfo = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&sslcert=%s&sslkey=%s&sslrootcert=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"), os.Getenv("CLIENT_CERT_PATH"), os.Getenv("CLIENT_KEY_PATH"), os.Getenv("SERVER_CA_PATH"))

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
