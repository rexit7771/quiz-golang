package controller

import (
	"net/http"
	"quiz-golang/database"
	"quiz-golang/repository"
	"quiz-golang/structs"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser structs.User
	err := c.BindJSON(&newUser)
	if err != nil {
		panic(err)
	}

	err = repository.Register(database.DbConnection, newUser)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "New User has been registered successfully")
}
