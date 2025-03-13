package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Variabel Konfigurasi Database
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "todo_db"
)

func ConnectDB() (*sql.DB, error) {

	//Membuat format string koneksi → Mirip kayak bikin alamat database supaya program bisa terhubung.
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//Coba buka koneksi ke PostgreSQL dengan informasi yang tadi dibuat.
	db, err := sql.Open("postgres", connStr)

	//Cek Apakah Ada Error
	if err != nil {
		return nil, err
	}
	//Mengembalikan Koneksi Database - Kalau berhasil, mengembalikan koneksi database supaya bisa dipakai di file lain.
	return db, nil
}
