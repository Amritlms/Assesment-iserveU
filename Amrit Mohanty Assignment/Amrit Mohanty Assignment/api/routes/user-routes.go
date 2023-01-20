package routes

import (
	"github.com/akhil/go-bookstore/api/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
}