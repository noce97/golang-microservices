package models

import (
	"fmt"
	"golang-microservices/2_ApplyingMVC/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User {
		123: {
			Id:        1,
			FirstName: "Pablo",
			LastName:  "Nocedal",
			Email:     "myemail@gmail.com",
		},
	}

	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError){
	log.Println("We are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v does not exist", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}