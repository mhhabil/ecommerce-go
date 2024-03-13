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
	FindOne(string) (*entity.User, error)
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
	_, errdb := service.db.Exec("insert into users(username,password,name) values ($1,$2,$3)", user.Username, pwd, user.Name)
	if errdb != nil {
		return user, errdb
	}
	return user, nil
}

func (service *userService) FindOne(username string) (*entity.User, error) {
	var user entity.User
	err := service.db.QueryRow("select username, name from users where username=$1", username).Scan(&user.Username, &user.Name)
	if err != nil {
		return &entity.User{}, err
	} else {
		return &user, nil
	}
}
