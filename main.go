package main

import (
	"net/http"
	"workplace/IMGo/controller"
	"workplace/IMGo/utils"
)

var (
	user = controller.UserControllerConstruct()
)

func main() {
	// handle request
	http.HandleFunc("/user/register", user.UserRegister)

	// handle static files
	utils.RegisterView()

	// server starts
	http.ListenAndServe(":8080", nil)
}


