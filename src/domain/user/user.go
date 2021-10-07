package user

import (
	"domain/skill"
	"gorm.io/gorm"
)

var user User

const USER_PREFIX string = "usr_"

type User struct {
	gorm.Model
	UserToken string        `gorm:"not null"` // 유저 토큰
	Email     string        `gorm:"not null"` // 이메일
	Username  string        `gorm:"not null"` // 유저 이름
	Password  string        `gorm:"not null"` // 유저 패스워드
	Skill     []skill.Skill // 보유 스킬들
}
