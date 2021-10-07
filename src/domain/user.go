package domain

import (
	"config"
	"gorm.io/gorm"
)

var user User

type User struct {
	gorm.Model
	UserToken string  `gorm:"not null"` // 유저 토큰
	Email     string  `gorm:"not null"` // 이메일
	Username  string  `gorm:"not null"` // 유저 이름
	Password  string  `gorm:"not null"` // 유저 패스워드
	Skill     []Skill // 보유 스킬들
}

func findByUserToken(userToken string) *gorm.DB {
	return config.DB.
		Where("user_token = ?", userToken).
		Find(&user)
}

func findByUserEmail(email string) *gorm.DB {
	return config.DB.
		Where("email = ?", email).
		Find(&user)
}

func findByUsername(username string) *gorm.DB {
	return config.DB.
		Where("username = ?", username).
		Find(&user)
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

func DeleteUser(userToken string) (err error) {
	config.DB.Where("user_token = ?", userToken).Delete(&user)
	return nil
}
