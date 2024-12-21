package controller

import (
	"net/http"
	"quiz-golang/database"
	"quiz-golang/repository"
	"quiz-golang/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBuku(c *gin.Context) {
	var result gin.H
	buku, err := repository.GetAllBuku(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": buku,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetBukuById(c *gin.Context) {
	var result gin.H
	var buku structs.Buku
	id, _ := strconv.Atoi(c.Param("id"))
	buku.ID = id
	data, err := repository.GetBukuById(database.DbConnection, buku)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": data,
		}
	}
	c.JSON(http.StatusOK, result)
}

func AddNewBuku(c *gin.Context) {
	var buku structs.Buku
	err := c.BindJSON(&buku)
	if err != nil {
		panic(err)
	}

	// * Validate Book's Released Year
	if buku.Release_year < 1980 {
		c.JSON(http.StatusBadRequest, "Book's Release Year only from 1980 - 2024")
	} else if buku.Release_year > 2024 {
		c.JSON(http.StatusBadRequest, "Book's Release Year only from 1980 - 2024")
	}

	err = repository.AddNewBuku(database.DbConnection, buku)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "New Book has been added succesfully")
}

func DeleteBuku(c *gin.Context) {
	var buku structs.Buku
	id, _ := strconv.Atoi(c.Param("id"))
	buku.ID = id
	_, err := repository.GetBukuById(database.DbConnection, buku)
	if err != nil {
		c.JSON(http.StatusNotFound, "Book with id "+c.Param("id")+" is not found")
		return
	} else {
		repository.DeleteBuku(database.DbConnection, buku)
		c.JSON(http.StatusOK, "buku with id "+c.Param("id")+" has been deleted successfully")
	}
}
