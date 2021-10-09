package route

import (
	skillController "controller/skill"
	userController "controller/user"
	"github.com/gin-gonic/gin"
)

func SetRoute() *gin.Engine {
	router := gin.Default()

	user := router.Group("/api/v1/users")
	{
		user.POST("/create", userController.CreateUser)
		//user.PATCH("/update/:userToken")
		//user.DELETE("/delete/:userToken")
	}

	skill := router.Group("/api/v1/skills")
	{
		skill.POST("/create", skillController.CreateSkill)
		//skill.PATCH("/update")
		//skill.DELETE("/delete")
	}

	return router
}
