package repository

import (
	"user-service/internal/entity/courier"
	"user-service/internal/entity/user"
)

type UserRepository interface {
	AddUser(req user.RegisterReq) (*user.UserRes, error)
	GetbyEmail(req user.LoginReq) (string, error)
	Getbyiduser(req user.UserRes) (*user.User, error)
	UpdateUser(req user.User) error
	DeleteUser(req user.UserRes) error

	AddCourier(courier courier.Courier) error
	GetbyCourierEmail(email string) (string, error)
	GetCourierByID(id int) (courier.Courier, error)
	UpdateCourier(courier courier.Courier) error
	DeleteCourier(id int) error
}
