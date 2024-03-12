package main

import (
	"ecommerce/controller"
	"ecommerce/database"
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

	authRoutes := server.Group("/auth")
	{
		authRoutes.POST("/register", func(ctx *gin.Context) {
			err := userController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "OK",
				})
			}
		})
	}

	server.Run(":8000")
}
