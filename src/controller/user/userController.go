package user

import (
	userService "domain/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var user userService.User

func CreateUser(c *gin.Context) {

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := userService.CreateUser(&user); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userToken := c.Params.ByName("userToken")

	if err := userService.DeleteUser(userToken); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"userToken" + userToken: "is deleted"})
}
