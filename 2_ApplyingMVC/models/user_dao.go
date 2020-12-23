package models

import (
	"fmt"
	"golang-microservices/2_ApplyingMVC/utils"
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
)

func GetUser(userId int64) (*User, *utils.ApplicationError){
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}