package main

import (
	"ecommerce/controller"
	"ecommerce/database"
	"ecommerce/middleware"
	"ecommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	var (
		userService    service.UserService       = service.New(db)
		userController controller.UserController = controller.New(
			userService,
		)
	)
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	routes := server.Group("/v1")
	productRoutes := routes.Group("/product")
	productRoutes.Use(middleware.AuthMiddleware())
	{
		productRoutes.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})
	}
	authRoutes := routes.Group("/user")
	{
		authRoutes.POST("/register", func(ctx *gin.Context) {
			code, err := userController.Save(ctx)
			if err != nil {
				ctx.JSON(code, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(code, gin.H{
					"message": "OK",
				})
			}
		})
		authRoutes.POST("/login", func(ctx *gin.Context) {
			res, code, err := userController.Login(ctx)
			if err != nil {
				ctx.JSON(code, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(code, res)
			}
		})
	}

	server.Run(":8000")
}
