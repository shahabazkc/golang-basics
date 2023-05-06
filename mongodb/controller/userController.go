package controller

import (
	"encoding/json"
	"fmt"
	"log"
	model "mongodb/models"
	"mongodb/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allUsers := services.GetAllUsers()
	json.NewEncoder(w).Encode(allUsers)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var user model.UserModel
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	services.CreateOneUser(user) // Call the GetAllUsers function of the userService instance

	json.NewEncoder(w).Encode(user)
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)

	type updateData struct {
		IsActive bool `json:"isActive"`
	}
	
	var data updateData

	if err := json.NewDecoder(r.Body).Decode(&data);err != nil{
		log.Fatal(err)
	}

	services.UpdateAUser(params["id"],data.IsActive)

	json.NewEncoder(w).Encode(params["id"])
}



func DeleteAUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	params := mux.Vars(r)

	services.DeleteOneUser(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}


func DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	deletedCount := services.DeleteAllUser()
	fmt.Println("Deleted count: ",deletedCount)
	json.NewEncoder(w).Encode("Deleted successfully.")
}




