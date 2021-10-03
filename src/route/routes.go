package route

import "github.com/gin-gonic/gin"

func UserRoute() {
	router := gin.Default()

	user := router.Group("/api/v1/users")
	{
		user.POST("/create")
		user.PATCH("/update")
		user.DELETE("/delete")
	}
}

func SkillRoute() {
	router := gin.Default()

	skill := router.Group("/api/v1/skills")
	{
		skill.POST("/create")
		skill.PATCH("/update")
		skill.DELETE("/delete")
	}
}
