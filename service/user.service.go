package service

import (
	"database/sql"
	"ecommerce/entity"
	"ecommerce/helper"
	"html"
	"net/http"
	"strings"

	"github.com/lib/pq"
)

type UserService interface {
	SaveUser(entity.User) (entity.User, int, error)
	FindUserByUsername(string) (*entity.User, error)
}

type userService struct {
	db    *sql.DB
	users []entity.User
}

func NewUserService(db *sql.DB) UserService {
	return &userService{
		db: db,
	}
}

func (service *userService) SaveUser(user entity.User) (entity.User, int, error) {

	pwd, err := helper.HashAndSalt(user.Password)
	if err != nil {
		return user, http.StatusInternalServerError, err
	}
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	_, errdb := service.db.Exec("insert into users(username,password,name) values ($1,$2,$3)", user.Username, pwd, user.Name)
	if errdb, ok := errdb.(*pq.Error); ok {
		if errdb.Code.Name() == "unique_violation" {
			return user, http.StatusConflict, errdb
		}
		return user, http.StatusInternalServerError, errdb
	}
	return user, http.StatusCreated, nil
}

func (service *userService) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := service.db.QueryRow("select username, name, password from users where username=$1", username).Scan(&user.Username, &user.Name, &user.Password)
	if err != nil {
		return &entity.User{}, err
	} else {
		return &user, nil
	}
}
