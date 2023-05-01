package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	var num int = 10
	fmt.Println("Value of num: ",num);
	fmt.Println("Address of num: ",&num);

	var numPtr = &num
	fmt.Println("Value of numPtr: ",*numPtr);

	*numPtr = *numPtr + 10
	fmt.Println("Value of num: ",num,numPtr);
}