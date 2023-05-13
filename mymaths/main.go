package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	// "math/rand"
	// "time"
)

func main(){
	fmt.Println("Welcome to maths in golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4.9

	fmt.Println("The sum is: ",mynumberOne + int(mynumberTwo))

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) +1 )
	// fmt.Println(crypto.)

	randomNum,err := rand.Int(rand.Reader,big.NewInt(5))

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(randomNum)
}