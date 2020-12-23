package controllers

import (
	"encoding/json"
	"golang-microservices/2_ApplyingMVC/services"
	"golang-microservices/2_ApplyingMVC/utils"
	"net/http"
	"strconv"
)

func GetUsers(response http.ResponseWriter, req *http.Request ){
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue,_ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		// Handle the error and return to the client
		jsonValue,_ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	// Return user to the client
	jsonValue,_ := json.Marshal(user)
	response.Write(jsonValue)
}
