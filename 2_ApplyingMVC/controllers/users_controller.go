package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-microservices/2_ApplyingMVC/services"
	"golang-microservices/2_ApplyingMVC/utils"
	"net/http"
	"strconv"
)

func GetUsers(c * gin.Context){
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.ErrorResponse(c, apiErr)
		return
	}


	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to the client
		utils.ErrorResponse(c, apiErr)
		return
	}

	// Return user to the client
	utils.Response(c, http.StatusOK, user)
}
