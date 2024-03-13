package helper

import (
	"ecommerce/entity"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error) {
	password := []byte(pwd)
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	salt, _ := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidatePassword(password string, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(pwd))
}

func GenerateToken(user entity.LoginRequest) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := entity.CustomClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
