package user

import (
	"common"
	"config"
	"gorm.io/gorm"
)

type UserService interface {
	FindByUserToken(userToken string)
	FindByUserEmail(email string)
	FindByUsername(username string)
	CreateUser(user *User)
	UpdateUser(user *User)
	DeleteUser(userToken string)
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
