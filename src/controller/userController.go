package controller

import (
	"domain/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var userService = user.UserService()

func CreateUser(c *gin.Context) {
	var user user.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := userService.CreateUser; err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, user)
}
