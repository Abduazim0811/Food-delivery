package userservice

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"user-service/internal/entity/courier"
	"user-service/internal/entity/user"
	"user-service/internal/infrastructura/redis"
	"user-service/internal/pkg/email"
	"user-service/internal/pkg/jwt"
	"user-service/internal/service"
	"user-service/userproto"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userproto.UnimplementedUserServiceServer
	service *service.UserService
	redis   *redis.RedisClient
}

func NewService(service *service.UserService, redis *redis.RedisClient) *Service {
	return &Service{service: service, redis: redis}
}

func (s *Service) Register(ctx context.Context, req *userproto.RegisterReq) (*userproto.RegisterRes, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	req.Password = string(bytes)

	code := 10000 + rand.Intn(90000)
	err = email.SendEmail(req.Email, email.SendClientCode(code, req.Firstname))
	if err != nil {
		log.Println("Error sending email to user:", err)
		return nil, fmt.Errorf("error sending email to user: %v", err)
	}

	userData := map[string]interface{}{
		"firstName": req.Firstname,
		"lastName":  req.Lastname,
		"phone":     req.Phone,
		"email":     req.Email,
		"password":  req.Password,
		"code":      code,
	}

	if s.redis == nil {
		log.Println("Redis client is not initialized")
		return nil, fmt.Errorf("redis client is not initialized")
	}

	err = s.redis.SetHash(req.Email, userData)
	if err != nil {
		log.Println("Failed to save user data in Redis:", err)
		return nil, fmt.Errorf("failed to save user data in Redis: %v", err)
	}

	return &userproto.RegisterRes{Message: "Verify code sent"}, nil
}

func (s *Service) VerifyCode(ctx context.Context, req *userproto.UserReq) (*userproto.UserRes, error) {
	res, err := s.redis.VerifyEmail(ctx, req.Email, req.Code)
	if err != nil {
		log.Println("Verify code error:", err)
		return nil, fmt.Errorf("verify code error: %v", err)
	}
	var userreq user.RegisterReq
	userreq.FirstName = res.Firstname
	userreq.LastName = res.Lastname
	userreq.Phone = res.Phone
	userreq.Email = res.Email
	userreq.Password = res.Password

	userres, err := s.service.Createuser(userreq)
	if err != nil {
		log.Println("Create user error:", err)
		return nil, fmt.Errorf("create user error: %v", err)
	}

	return &userproto.UserRes{Id: int32(userres.ID)}, nil
}

func (s *Service) Login(ctx context.Context, req *userproto.LoginReq) (*userproto.LoginRes, error) {
	password, err := s.service.GetByEmail(user.LoginReq{Email: req.Email})
	if err != nil {
		log.Println("Login error:", err)
		return nil, fmt.Errorf("login error: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	token, err := jwt.GenerateToken(req.Email, "user")
	if err != nil {
		return nil, err
	}
	_, err = jwt.CheckToken(token)
	if err != nil {
		log.Println("token error: ", err)
		return nil, fmt.Errorf("token error: %v", err)
	}
	return &userproto.LoginRes{Token: token}, nil
}

func (s *Service) GetbyIdUser(ctx context.Context, req *userproto.UserRes) (*userproto.User, error) {
	res, err := s.service.GetByiduser(user.UserRes{ID: int(req.Id)})
	if err != nil {
		log.Println("Error fetching user by ID:", err)
		return nil, fmt.Errorf("error fetching user by ID: %v", err)
	}

	return &userproto.User{
		Id:        int32(res.ID),
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		Phone:     res.Phone,
		Email:     res.Email,
		Password:  res.Password,
	}, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *userproto.User) (*userproto.RegisterRes, error) {
	var users user.User
	users.ID = int(req.Id)
	users.FirstName = req.Firstname
	users.LastName = req.Lastname
	users.Phone = req.Phone
	users.Email = req.Email
	users.Password = req.Password

	err := s.service.Updateuser(users)
	if err != nil {
		log.Println("Error updating user:", err)
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return &userproto.RegisterRes{Message: "User updated"}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *userproto.UserRes) (*userproto.RegisterRes, error) {
	err := s.service.Deleteuser(user.UserRes{ID: int(req.Id)})
	if err != nil {
		log.Println("Error deleting user:", err)
		return nil, fmt.Errorf("error deleting user: %v", err)
	}

	return &userproto.RegisterRes{Message: "User deleted"}, nil
}

func (s *Service) RegisterCourier(ctx context.Context,req *userproto.RegisterCourierRequest) (*userproto.RegisterCourierResponse, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	req.Password = string(bytes)

	code := 10000 + rand.Intn(90000)
	err = email.SendEmail(req.Email, email.SendClientCode(code, req.Name))
	if err != nil {
		log.Println("Error sending email to user:", err)
		return nil, fmt.Errorf("error sending email to user: %v", err)
	}
	userData := map[string]interface{}{
		"name": req.Name,
		"phoneNumber": req.PhoneNumber,
		"email":     req.Email,
		"password":  req.Password,
		"code":      code, 
	}

	if s.redis == nil {
		log.Println("Redis client is not initialized")
		return nil, fmt.Errorf("redis client is not initialized")
	}

	err = s.redis.SetHash(req.Email, userData)
	if err != nil {
		log.Println("Failed to save user data in Redis:", err)
		return nil, fmt.Errorf("failed to save user data in Redis: %v", err)
	}

	return &userproto.RegisterCourierResponse{Message: "verify code"}, nil
}

func (s *Service) VerifyCodeCourier(ctx context.Context, req *userproto.UserReq) (*userproto.UserRes, error) {
    res, err := s.redis.VerifyEmailCourier(ctx, req.Email, req.Code)
    if err != nil {
        log.Println("Verify code error:", err)
        return nil, fmt.Errorf("verify code error: %v", err)
    }

    courierData := courier.Courier{
        Name:        res.Name,
        Email:       res.Email,
        PhoneNumber: res.PhoneNumber,
        PasswordHash:    res.Password,
    }

    err = s.service.Createcourier(courierData)
	if err != nil {
		log.Println("error create courier:", err)
		return nil, fmt.Errorf("error create courier: %v", err)
	}

    return &userproto.UserRes{Id: int32(courierData.ID)}, nil
}

func (s *Service) LoginCourier(ctx context.Context, req *userproto.LoginCourierRequest) (*userproto.LoginCourierResponse, error) {
    password, err := s.service.Logincourier(req.Email)
    if err != nil {
        log.Println("Login error:", err)
        return nil, fmt.Errorf("login error: %v", err)
    }

    err = bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password))
    if err != nil {
        return nil, fmt.Errorf("incorrect password")
    }

    token, err := jwt.GenerateToken(req.Email, "courier")
    if err != nil {
        return nil, err
    }

    return &userproto.LoginCourierResponse{Token: token, Message: "Login successful"}, nil
}

func (s *Service) UpdateCourier(ctx context.Context, req *userproto.Courier) (*userproto.RegisterCourierResponse, error) {
    var hashedPassword string
    if req.Password != "" {
        bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        if err != nil {
            log.Println("Error hashing password:", err)
            return nil, fmt.Errorf("error hashing password: %v", err)
        }
        hashedPassword = string(bytes)
    }

    courier := courier.Courier{
        ID:           int(req.Id),
        Name:         req.Name,
        Email:        req.Email,
        PhoneNumber:  req.PhoneNumber,
        PasswordHash: hashedPassword,
        // Update other fields if necessary
    }

    err := s.service.Updatecourier(courier)
    if err != nil {
        log.Println("Error updating courier:", err)
        return nil, fmt.Errorf("error updating courier: %v", err)
    }

    return &userproto.RegisterCourierResponse{Message: "Courier updated successfully"}, nil
}


func (s *Service) DeleteCourier(ctx context.Context, req *userproto.UserRes) (*userproto.RegisterCourierResponse, error) {
    err := s.service.Deletecourier(int(req.Id))
    if err != nil {
        log.Println("Error deleting courier:", err)
        return nil, fmt.Errorf("error deleting courier: %v", err)
    }

    return &userproto.RegisterCourierResponse{Message: "Courier deleted successfully"}, nil
}