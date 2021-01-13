package services

import (
	"golang-microservices/2_ApplyingMVC/models"
	"golang-microservices/2_ApplyingMVC/utils"
)

type usersService struct {
}

var(
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*models.User, *utils.ApplicationError) {
	user, err := models.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}