package userhandler

import (
	_ "api-gateway/docs"
	producer "api-gateway/internal/infrastructura/kafka"
	"api-gateway/internal/protos/userproto"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ClientUser userproto.UserServiceClient
}

// @title Food Delivery
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @securityDefinitions.apikey Bearer
// @in 				header
// @name Authorization
// @description Enter the token in the format `Bearer {token}`
// @host localhost:7777
// @BasePath /

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.RegisterReq true "User request body"
// @Success 200 {object} userproto.RegisterRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /users/register [post]
func (u *UserHandler) Register(c *gin.Context) {
	var req userproto.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.Register(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	byted, err := json.Marshal(&req)
	if err != nil {
		log.Println(err)
	}
	if err := producer.Producer("create", byted); err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Verify user code
// @Description Verify the user's code for registration or password reset
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.UserReq true "Verification request"
// @Success 200 {object} userproto.UserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /users/verify-code [post]
func (u *UserHandler) VerifyCode(c *gin.Context) {
	var req userproto.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.VerifyCode(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) Login(c *gin.Context) {
	var req userproto.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get user by ID
// @Description Retrieve user details by user ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} userproto.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /users/{id} [get]
func (u *UserHandler) GetbyIdUser(c *gin.Context) {
	id := c.Param("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	req := &userproto.UserRes{Id: int32(user_id)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.GetbyIdUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update user details
// @Description Update the details of an existing user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.User true "User update request"
// @Success 200 {object} userproto.RegisterRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /users/ [put]
func (u *UserHandler) UpdateUser(c *gin.Context) {
	var req userproto.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.UpdateUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	byted, err := json.Marshal(&req)
	if err != nil {
		log.Println(err)
	}
	if err := producer.Producer("update", byted); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Delete user by ID
// @Description Delete an existing user by user ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} userproto.RegisterRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /users/{id} [delete]
func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	req := &userproto.UserRes{Id: int32(user_id)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.DeleteUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	byted, err := json.Marshal(&req)
	if err != nil {
		log.Println(err)
	}
	if err := producer.Producer("delete", byted); err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Register a new courier
// @Description Create a new courier with the given details
// @Tags courier
// @Accept json
// @Produce json
// @Param courier body userproto.RegisterCourierRequest true "Courier registration request"
// @Success 200 {object} userproto.RegisterCourierResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /couriers/register [post]
func (u *UserHandler) RegisterCourier(c *gin.Context) {
	var req userproto.RegisterCourierRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.RegisterCourier(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Verify courier code
// @Description Verify the courier's code for registration or password reset
// @Tags courier
// @Accept json
// @Produce json
// @Param user body userproto.UserReq true "Verification request"
// @Success 200 {object} userproto.UserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /couriers/verify-code [post]
func (u *UserHandler) VerifyCodeCourier(c *gin.Context) {
	var req userproto.UserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.VerifyCodeCourier(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Login courier
// @Description Authenticate a courier and return an access token
// @Tags courier
// @Accept json
// @Produce json
// @Param login body userproto.LoginCourierRequest true "Courier login request"
// @Success 200 {object} userproto.LoginCourierResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /couriers/login [post]
func (u *UserHandler) LoginCourier(c *gin.Context) {
	var req userproto.LoginCourierRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.LoginCourier(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update courier details
// @Description Update the details of an existing courier
// @Tags courier
// @Accept json
// @Produce json
// @Param courier body userproto.Courier true "Courier update request"
// @Success 200 {object} userproto.Courier
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /couriers/ [put]
func (u *UserHandler) UpdateCourier(c *gin.Context) {
	var req userproto.Courier

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.UpdateCourier(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete courier by ID
// @Description Delete an existing courier by courier ID
// @Tags courier
// @Accept json
// @Produce json
// @Param id path int true "Courier ID"
// @Success 200 {object} userproto.RegisterCourierResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /couriers/{id} [delete]
func (u *UserHandler) DeleteCourier(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	req := &userproto.UserRes{Id: int32(Id)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.DeleteCourier(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
