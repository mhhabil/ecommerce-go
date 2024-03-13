package helper

import (
	"ecommerce/entity"
	"os"
	"strconv"
	"time"

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
	var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Minute
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, jwt.MapClaims{
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"eat":      time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
	})

	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
