package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-golang/controller"
	"quiz-golang/database"
	"quiz-golang/middlewares"

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
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
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
	router.POST("/api/register", controller.Register)
	router.Use(middlewares.Auth())
	router.GET("/api/categories", controller.GetAllKategori)
	router.POST("/api/categories", controller.AddNewKategori)
	router.GET("/api/categories/:id/books", controller.GetBukuByKategoriId)
	router.GET("/api/categories/:id", controller.GetKategoriById)
	router.DELETE("/api/categories/:id", controller.DeleteKategori)

	router.GET("/api/books", controller.GetAllBuku)
	router.POST("/api/books", controller.AddNewBuku)
	router.GET("/api/books/:id", controller.GetBukuById)
	router.DELETE("/api/books/:id", controller.DeleteBuku)

	router.Run(":" + os.Getenv("PORT"))
}
