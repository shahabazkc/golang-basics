package router

import (
	"mongodb/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/create",controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/users",controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}",controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/delete/{id}",controller.DeleteAUser).Methods("DELETE")
	router.HandleFunc("/api/user/delete-all",controller.DeleteAllUser).Methods("DELETE")

	return router
}