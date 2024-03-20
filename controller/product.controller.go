package controller

import (
	"ecommerce/entity"
	"ecommerce/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController interface {
	SaveProduct(ctx *gin.Context) (int, error)
	UpdateProduct(ctx *gin.Context) (int, error)
	// Delete(ctx *gin.Context) (int, error)
	// FindAll(ctx *gin.Context) ([]entity.Product, int, error)
	// FindById(ctx *gin.Context) (entity.Product, int, error)
}

type productController struct {
	service service.ProductService
}

var validate *validator.Validate

func NewProductController(service service.ProductService) ProductController {
	validate = validator.New()
	return &productController{
		service: service,
	}
}

func (c *productController) SaveProduct(ctx *gin.Context) (int, error) {
	var product entity.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return http.StatusBadRequest, err
	}

	errVal := validate.Struct(product)
	if errVal != nil {
		fmt.Println("err", errVal)
		return http.StatusBadRequest, errVal
	}
	_, code, errdb := c.service.SaveProduct(product)
	if errdb != nil {
		return code, errdb
	}
	return code, nil
}

func (c *productController) UpdateProduct(ctx *gin.Context) (int, error) {
	var product entity.Product
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return http.StatusBadRequest, err
	}

	errVal := validate.Struct(product)
	if errVal != nil {
		return http.StatusBadRequest, errVal
	}
	_, code, errdb := c.service.UpdateProduct(product, id)
	if errdb != nil {
		return code, errdb
	}
	return code, nil
}
