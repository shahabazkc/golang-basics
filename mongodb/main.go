package main

import (
	"fmt"
	"log"
	"mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("mongo db")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Server is started on port 4000")
}