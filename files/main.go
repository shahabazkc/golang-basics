package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcome to files")
	createAndWriteFile("Hi, there!");
	readFile("./myfile.txt")
}

func checkNilErr(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		panic(err)
	}
}

func readFile(filename string){
	databytes,err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Data as bytes: ", string(databytes))
}

func createAndWriteFile(content string)bool{
	file,err:= os.Create("./myfile.txt")

	checkNilErr(err)

	length,err := file.WriteString(content)

	checkNilErr(err)

	fmt.Println(length, "bytes written successfully");
	defer file.Close();

	return true;
}