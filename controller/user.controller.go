package controller

import (
	"ecommerce/entity"
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindOne() entity.User
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

func (c *controller) FindOne() entity.User {
	return c.service.FindOne()
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
