package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-golang/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s 
	password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	database.DBMigrate(DB)
}
