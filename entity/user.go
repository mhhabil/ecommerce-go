package entity

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username" binding:"min=5,max=15,required"`
	Name     string `json:"name" binding:"min=5,max=50,required"`
	Password string `json:"password" binding:"min=5,max=15,required"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseData struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type LoginResponse struct {
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
