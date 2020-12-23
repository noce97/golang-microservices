package app

import (
	"golang-microservices/2_ApplyingMVC/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
