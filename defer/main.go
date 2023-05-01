package main

import "fmt"

func main(){
	defer fmt.Println("World")
	defer fmt.Println("Bye")
	fmt.Println("Welcome to defer")
	fmt.Println("Hello");
	myDef();
};

func myDef(){
	fmt.Println("myDef")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}
}