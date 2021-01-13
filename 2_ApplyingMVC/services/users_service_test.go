package services

import (
	"github.com/stretchr/testify/assert"
	"golang-microservices/2_ApplyingMVC/models"
	"golang-microservices/2_ApplyingMVC/utils"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64)(*models.User, *utils.ApplicationError)
)

func init(){
	models.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {
}

func (m * usersDaoMock) GetUser(userId int64) (*models.User, *utils.ApplicationError){
	return getUserFunction(userId)
}

func TestUsersService_GetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*models.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "User 0 does not exist",
			StatusCode: http.StatusNotFound,
		}
	}

	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 does not exist", err.Message)
}

func TestUsersService_GetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*models.User, *utils.ApplicationError) {
		return &models.User{
			Id:        1,
		}, nil
	}

	user, err := UsersService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 1, user.Id)
}