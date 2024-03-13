package middleware

import (
	"ecommerce/entity"
	"ecommerce/helper"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load .env",
			})
			ctx.Abort()
			return
		}
		var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
		var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
		tokenString := helper.ExtractToken(ctx)

		//parse read expires
		token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SECRET_KEY, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		_, ok := token.Claims.(*entity.CustomClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
