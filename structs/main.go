package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs")
	User1 := User{"John", "Doe", 30, "sh", true}
	fmt.Printf("User1: %+v\n", User1);
	fmt.Printf("Name: %v %v and email is: %v", User1.FirstName, User1.LastName,User1.Email);
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	IsActive  bool
}
