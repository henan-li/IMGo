package main

import (
	"net/http"
	"workplace/IMGo/controller"
	"workplace/IMGo/utils"
)

var (
	user = controller.UserControllerConstruct()
	contact = controller.ContactControllerConstruct()
)

func main() {
	// user api
	http.HandleFunc("/user/login", user.UserLogin)
	http.HandleFunc("/user/register", user.UserRegister)

	// contact api
	http.HandleFunc("/contact/addfriend", contact.AddFriends)
	http.HandleFunc("/contact/joingroup", contact.JoinGroup)
	http.HandleFunc("/contact/loadfriend", contact.LoadFriend)
	http.HandleFunc("/contact/loadgroup", contact.LoadGroup)
	http.HandleFunc("/contact/createcommunity", contact.CreateGroup)

	// chat
	http.HandleFunc("/chat", controller.Chat)

	// handle static files
	utils.RegisterView()

	// server starts
	http.ListenAndServe(":8080", nil)
}


