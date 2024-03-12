package main

import (
	"ecommerce/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.Run(":8000")
}
