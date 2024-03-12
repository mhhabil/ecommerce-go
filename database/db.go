package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occured on .env file", err)
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("PostgreSQL ready!!!")
	}

	return db, err
}
