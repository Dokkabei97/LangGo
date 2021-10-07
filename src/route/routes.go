package route

import (
	"controller"
	"github.com/gin-gonic/gin"
)

func UserRoute() *gin.Engine {
	router := gin.Default()

	user := router.Group("/api/v1/users")
	{
		user.POST("/create", controller.CreateUser)
		user.PATCH("/update/:userToken")
		user.DELETE("/delete/:userToken")
	}

	return router
}

func SkillRoute() *gin.Engine {
	router := gin.Default()

	skill := router.Group("/api/v1/skills")
	{
		skill.POST("/create")
		skill.PATCH("/update")
		skill.DELETE("/delete")
	}

	return router
}
