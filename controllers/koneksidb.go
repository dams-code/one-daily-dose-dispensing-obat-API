package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"one-daily-dose-dispensing-obat-api/migrations"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var DBSqlConn *sql.DB

func KoneksiDB() (*sql.DB, error) {

	err := godotenv.Load("config/.env")

	if err != nil {
		log.Println("File config/.env tidak ditemukan")
	}

	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err := sql.Open("postgres", psqlConn)

	if err != nil {
		log.Fatalf("Gagal terhubung ke postgres %v", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke postgres %v", err.Error())
	}

	if err = migrations.MigrasiDataObat(db); err != nil {
		log.Fatalf("Error Migrasi data obat : %v", err.Error())
	}

	fmt.Println("Berhasil terhubung ke database postgres obat.")

	return db, nil
}
