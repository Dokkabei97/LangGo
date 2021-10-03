package main

import (
	"config"
	"domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(config.DBUrl(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패했습니다!")
	}

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Skill{})

	db.Create(&domain.User{Username: "test", Skill: nil})
	db.Create(&domain.Skill{Lang: "Java", Tier: 1, UserID: 1})
	db.Create(&domain.Skill{Lang: "Go", Tier: 2, UserID: 1})
	db.Create(&domain.Skill{Lang: "Python", Tier: 2, UserID: 1})
	db.Create(&domain.Skill{Lang: "JavaScript", Tier: 2, UserID: 1})
}
