package skill

import (
	skillService "domain/skill"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateSkill(c *gin.Context) {
	var skill skillService.Skill

	if err := c.BindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := skillService.CreateSkill(&skill); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, skill)
}
