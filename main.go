package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-golang/controller"
	"quiz-golang/database"

	"github.com/gin-gonic/gin"
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
	// database.UnmigrateDB(DB)

	router := gin.Default()
	router.GET("/api/categories", controller.GetAllKategori)
	router.POST("/api/categories", controller.AddNewKategori)
	router.GET("/api/categories/:id/books", controller.GetBukuByKategoriId)
	router.GET("/api/categories/:id", controller.GetKategoriById)
	router.DELETE("/api/categories/:id", controller.DeleteKategori)

	router.Run(":" + os.Getenv("PORT"))
}
