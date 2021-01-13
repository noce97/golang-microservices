package services

import (
	"golang-microservices/2_ApplyingMVC/models"
	"golang-microservices/2_ApplyingMVC/utils"
	"net/http"
)

type itemsService struct{

}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemId string) (*models.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me please",
		StatusCode: http.StatusInternalServerError,
	}
}