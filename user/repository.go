package user

import "github.com/tenahubapi/entity"

type UserRepository interface {
	Users(role string)([]entity.User, []error)
	User(user *entity.User)(*entity.User, []error)
	UserByID(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User)(*entity.User, []error)
	DeleteUser(id uint)(*entity.User, []error)
	StoreUser(user *entity.User)(*entity.User, []error)
}