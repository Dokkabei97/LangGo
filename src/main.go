package main

import (
	"config"
	"domain/skill"
	"domain/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"route"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DBUrl(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		panic("DB 연결에 실패했습니다!")
	}

	err = config.DB.AutoMigrate(&user.User{}, &skill.Skill{})

	r := route.SetRoute()
	r.Use(gin.Logger())

	err = r.Run()
}
