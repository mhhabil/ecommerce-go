package controller

import (
	"ecommerce/entity"
	"ecommerce/helper"
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context) (entity.LoginResponse, error)
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) Login(ctx *gin.Context) (entity.LoginResponse, error) {
	var request entity.LoginRequest
	var response entity.LoginResponse
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		return entity.LoginResponse{}, err
	}
	user, err := c.service.FindOne(request.Username)
	if err != nil {
		return entity.LoginResponse{}, err
	}

	token, err := helper.GenerateToken(request)
	if err != nil {
		return entity.LoginResponse{}, err
	}

	response = entity.LoginResponse{
		Message: "OK",
		Data: entity.LoginResponseData{
			Username:    user.Username,
			Name:        user.Name,
			AccessToken: token,
		},
	}
	return response, nil
}

func (c *controller) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	_, errdb := c.service.Save(user)
	if errdb != nil {
		return errdb
	}
	return nil
}
