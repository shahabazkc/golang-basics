package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch")
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	fmt.Println("Number is: ", num)
	switch num {
	case 1:
		fmt.Println("Number is 1")
		fallthrough
	case 2:
		fmt.Println("Number is 2")
		fallthrough
	case 3:
		fmt.Println("Number is 3")
	default:
		fmt.Println("Number is not 1,2 or 3")
	}

}
