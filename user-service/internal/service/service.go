package service

import (
	"user-service/internal/entity/courier"
	"user-service/internal/entity/user"
	"user-service/internal/infrastructura/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Createuser(req user.RegisterReq) (*user.UserRes, error) {
	return u.repo.AddUser(req)
}

func (u *UserService) GetByEmail(req user.LoginReq) (string, error) {
	return u.repo.GetbyEmail(req)
}

func (u *UserService) GetByiduser(req user.UserRes) (*user.User, error) {
	return u.repo.Getbyiduser(req)
}

func (u *UserService) Updateuser(req user.User) error {
	return u.repo.UpdateUser(req)
}

func (u *UserService) Deleteuser(req user.UserRes) error {
	return u.repo.DeleteUser(req)
}

func (u *UserService) Createcourier(courier courier.Courier)error{
	return u.repo.AddCourier(courier)
}

func (u *UserService) Logincourier(email string)(string, error){
	return u.repo.GetbyCourierEmail(email)
}

func (u *UserService) Getcourierbyid(id int)(courier.Courier, error){
	return u.repo.GetCourierByID(id)
}

func (u *UserService) Updatecourier(courier courier.Courier)error{
	return u.repo.UpdateCourier(courier)
}

func (u *UserService) Deletecourier(id int)error{
	return u.repo.DeleteCourier(id)
}