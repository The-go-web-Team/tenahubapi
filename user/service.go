package user

import "github.com/TenaHub/api/entity"

type UserService interface {
	Users(role string) ([]entity.User, []error)
	User(user *entity.User) (*entity.User, []error)
	UserByID(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
}