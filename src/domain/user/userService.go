package user

type UserService interface {
	FindByUserToken(userToken string)
	FindByUserEmail(email string)
	FindByUsername(username string)
	CreateUser(user *User)
	UpdateUser(user *User)
	DeleteUser(userToken string)
}
