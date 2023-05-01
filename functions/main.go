package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions");
	fmt.Println(greet("John")); 
	fmt.Println(sum(1,2));

	var result= sum(1,5);
	fmt.Println(result);
	sumValues,str := sumOfValues(1,2,3,4,5);
	fmt.Println(sumValues,str);
}

func greet(name string)string {
	return "Hello " + name
};

func sumOfValues(values ...int)(int,string){
	sum := 0
	for _,i := range values{
		sum += i
	}
	return sum, "Success"
}


func sum(a int, b int)int{
	return a + b
}