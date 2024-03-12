package service

import (
	"database/sql"
	"ecommerce/entity"
	"ecommerce/helper"
	"html"
	"strings"
)

type UserService interface {
	Save(entity.User) (entity.User, error)
	FindOne() entity.User
}

type userService struct {
	db    *sql.DB
	users []entity.User
}

func New(db *sql.DB) UserService {
	return &userService{
		db: db,
	}
}

func (service *userService) Save(user entity.User) (entity.User, error) {

	pwd, err := helper.HashAndSalt(user.Password)
	if err != nil {
		return user, err
	}
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	_, errdb := service.db.Exec("insert into users(username,password,email) values ($1,$2,$3)", user.Username, pwd, user.Email)
	if errdb != nil {
		return user, errdb
	}
	return user, nil
}

func (service *userService) FindOne() entity.User {
	return service.users[0]
}
