package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func Reset(db *sql.DB) {

	var err error

	// Buat perintah SQL untuk mereset hit_chance
	query := "UPDATE auth SET hit_chance = $1"
	_, err = db.Exec(query, 1000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hit chance berhasil direset.")

	// Menampilkan waktu reset
	fmt.Printf("Waktu reset: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}
