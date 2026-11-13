package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectionDB() {

	// 1 load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2 mengambil valaue yang ada pada env
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	// 3 menyambungkan ke database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to the database")

	DB = db
}
