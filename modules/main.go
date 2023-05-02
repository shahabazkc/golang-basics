package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to mymodules")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	http.ListenAndServe(":4000",r)

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter(){
	fmt.Println("Hello from greeter")
}

func serverHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to shahabaz tutorial</h1>"))
}