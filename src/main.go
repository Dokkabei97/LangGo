package main

import (
	"domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패했습니다!")
	}

	db.AutoMigrate(&domain.Skill{})

	db.Create(&domain.Skill{Lang: "Java", Tier: 3})

	var skill domain.Skill
	db.First(&skill, 1)
	db.First(&skill, "lang = ?", "Java")

	db.Model(&skill).Update("Tier", 1)

	db.Delete(&skill, 1)
}
