package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitish-126/phonebook/pkg/models"
	"github.com/nitish-126/phonebook/pkg/utils"
)

// var newUser models.User

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Users(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	newUsers := models.GetUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UserbyName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	userDetails := models.GetUserByName(name)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UserbyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userDetails := models.GetUserById(id)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UserbyMobileNo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["mobileno"]
	userDetails := models.GetUserBymobileNo(num)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddUsers(w http.ResponseWriter, r *http.Request) {
	addUser := &models.User{}
	utils.ParseBody(r, addUser)
	b := addUser.CreateUser()
	json.Marshal(b)
	// res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func Deleteuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	models.DeleteUser(name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteuserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	models.DeleteUserById(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func Updateuser(w http.ResponseWriter, r *http.Request) {
	var UpdateUser = &models.User{}
	utils.ParseBody(r, UpdateUser)
	vars := mux.Vars(r)
	id := vars["id"]
	UpdateUser.UpdateUser(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
