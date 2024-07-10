package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitish-126/phonebook/pkg/controllers"
)

var CreatedContactRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Users).Methods("GET")
	router.HandleFunc("/username/{name}", controllers.UserbyName).Methods("GET")
	router.HandleFunc("/userid/{id}", controllers.UserbyId).Methods("GET")
	router.HandleFunc("/usernum/{mobileno}", controllers.UserbyMobileNo).Methods("GET")
	router.HandleFunc("/add-user", controllers.AddUsers).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.Updateuser).Methods("PUT")
	router.HandleFunc("/delete-name/{name}", controllers.Deleteuser).Methods("DELETE")
	router.HandleFunc("/delete-id/{id}", controllers.DeleteuserById).Methods("DELETE")
}
