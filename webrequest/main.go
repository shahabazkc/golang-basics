package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"


func main(){
	fmt.Println("Welcome to webrequest")

	res,err:=http.Get(url)

	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", res)

	databytes,err:= ioutil.ReadAll(res.Body)

	checkNilErr(err)
	fmt.Println("Data as string: ", string(databytes))

	defer res.Body.Close()
}


func checkNilErr(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		panic(err)
	}
}