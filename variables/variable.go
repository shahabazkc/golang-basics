package main

import "fmt"

func main() {
	var username string = "Shahabaz kc"
	fmt.Printf("Username is %T \n", username)
	fmt.Println("Hello, World!")
	// initial variable
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println(c)

	// float
	var d float64 = 10.52
	var e float64 = 20.52
	var f float64 = d + e
	fmt.Println(f)

	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn)
	fmt.Printf(("Is logged in %T \n"), isLoggedIn)

	var smallVal uint64= 120000000000000000
	fmt.Println(smallVal)

	var anotherVal int
	fmt.Println(anotherVal);

	var website = "shahabaz kc";
	fmt.Println(website);

	numberOfUser:= 1000;
	fmt.Println(numberOfUser);

}
