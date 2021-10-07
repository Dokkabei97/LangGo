package main

import (
	"common"
	"config"
	"domain/skill"
	"domain/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(config.DBUrl(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패했습니다!")
	}

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&skill.Skill{})

	pwd := "test123!"
	pwdCry, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	db.Create(&user.User{
		UserToken: common.RandomCharacterWithPrefix(user.USER_PREFIX),
		Email:     "test@test.com",
		Password:  string(pwdCry),
		Username:  "test",
		Skill:     nil,
	})

	db.Create(&skill.Skill{
		SkillToken: common.RandomCharacterWithPrefix(skill.Skill_PREFIX),
		Lang:       "Java", Tier: 1, UserID: 1,
	})
	db.Create(&skill.Skill{
		SkillToken: common.RandomCharacterWithPrefix(skill.Skill_PREFIX),
		Lang:       "Go", Tier: 2, UserID: 1,
	})
	db.Create(&skill.Skill{
		SkillToken: common.RandomCharacterWithPrefix(skill.Skill_PREFIX),
		Lang:       "Python", Tier: 2, UserID: 1,
	})
	db.Create(&skill.Skill{
		SkillToken: common.RandomCharacterWithPrefix(skill.Skill_PREFIX),
		Lang:       "JavaScript", Tier: 2, UserID: 1,
	})

}
