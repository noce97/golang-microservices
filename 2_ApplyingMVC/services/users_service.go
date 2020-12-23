package services

import (
	"golang-microservices/2_ApplyingMVC/models"
	"golang-microservices/2_ApplyingMVC/utils"
)

func GetUser(userId int64) (*models.User, *utils.ApplicationError) {
	return  models.GetUser(userId)
}