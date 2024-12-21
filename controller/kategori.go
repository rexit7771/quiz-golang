package controller

import (
	"net/http"
	"quiz-golang/database"
	"quiz-golang/repository"
	"quiz-golang/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllKategori(c *gin.Context) {
	var result gin.H
	kategori, err := repository.GetAllKategori(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": kategori,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetKategoriById(c *gin.Context) {
	var result gin.H
	var kategori structs.Kategori
	id, _ := strconv.Atoi(c.Param("id"))
	kategori.ID = id
	data, err := repository.GetKategoriById(database.DbConnection, kategori)
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

func GetBukuByKategoriId(c *gin.Context) {
	var result gin.H
	var kategori structs.Kategori
	id, _ := strconv.Atoi(c.Param("id"))
	kategori.ID = id
	buku, err := repository.GetBukuByKategoriId(database.DbConnection, kategori)
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

func AddNewKategori(c *gin.Context) {
	var newKategori structs.Kategori
	err := c.BindJSON(&newKategori)
	if err != nil {
		panic(err)
	}

	err = repository.AddNewKategori(database.DbConnection, newKategori)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, "New Category has been added successfully")
}

func DeleteKategori(c *gin.Context) {
	var kategori structs.Kategori
	id, _ := strconv.Atoi(c.Param("id"))
	kategori.ID = id
	err := repository.DeleteKategori(database.DbConnection, kategori)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, kategori)
}
