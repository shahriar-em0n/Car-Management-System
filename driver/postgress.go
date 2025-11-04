package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	db *sql.DB
)

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Waiting for the database Start up...")
	time.Sleep(5 * time.Second)

	db, err := sql.Open("postgress", connStr)
	if err != nil {
		log.Fatalf("Error Opeing database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error Connection to the database: %v", err)
	}

	fmt.Println("SuccesFully Connected to the database")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatalf("Error Closing The Database: %v", err)
	}
}
