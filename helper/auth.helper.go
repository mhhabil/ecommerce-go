package helper

import (
	"os"
	"strconv"

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
