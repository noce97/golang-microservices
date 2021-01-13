package app

import (
	"golang-microservices/2_ApplyingMVC/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUsers)
}