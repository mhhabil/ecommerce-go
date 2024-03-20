package controller

import (
	"ecommerce/entity"
	"ecommerce/helper"
	"ecommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context) (entity.LoginResponse, int, error)
	SaveUser(ctx *gin.Context) (int, error)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Login(ctx *gin.Context) (entity.LoginResponse, int, error) {
	var request entity.LoginRequest
	var response entity.LoginResponse
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		return entity.LoginResponse{}, http.StatusBadRequest, err
	}
	user, err := c.service.FindUserByUsername(request.Username)
	if err != nil {
		return entity.LoginResponse{}, http.StatusNotFound, err
	}

	errPassword := helper.ValidatePassword(user.Password, request.Password)
	if errPassword != nil {
		return entity.LoginResponse{}, http.StatusBadRequest, errPassword
	}

	token, err := helper.GenerateToken(request)
	if err != nil {
		return entity.LoginResponse{}, http.StatusInternalServerError, err
	}

	response = entity.LoginResponse{
		Message: "OK",
		Data: entity.LoginResponseData{
			Username:    user.Username,
			Name:        user.Name,
			AccessToken: token,
		},
	}
	return response, 200, nil
}

func (c *userController) SaveUser(ctx *gin.Context) (int, error) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return http.StatusBadRequest, err
	}
	_, code, errdb := c.service.SaveUser(user)
	if errdb != nil {
		return code, errdb
	}
	return code, nil
}
