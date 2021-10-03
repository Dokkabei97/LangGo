package domain

import (
	"config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  // 유저 이름
	Skill    []Skill // 보유 스킬들
}

func findById(id uint) {
	// todo 아이디 조회
}

func findByUsername(username string) {
	// todo 유저이름 검색
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User) (err error) {
	// todo 유저 업데이트
	return nil
}

func DeleteUser(user *User) (err error) {
	// todo 유저 삭제
	return nil
}
