package user

import (
	"common"
	"config"
	"domain/skill"
	"golang.org/x/crypto/bcrypt"
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

func FindByUserToken(userToken string) *gorm.DB {
	return config.DB.
		Where("user_token = ?", userToken).
		Find(&user)
}

func FindByUserEmail(email string) *gorm.DB {
	return config.DB.
		Where("email = ?", email).
		Find(&user)
}

func FindByUsername(username string) *gorm.DB {
	return config.DB.
		Where("username = ?", username).
		Find(&user)
}

func CreateUser(user *User) (err error) {
	user.UserToken = common.RandomCharacterWithPrefix(USER_PREFIX)
	encryptPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(encryptPwd)
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
