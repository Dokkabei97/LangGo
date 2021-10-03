package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  // 유저 이름
	Skill    []Skill // 보유 스킬들
}
