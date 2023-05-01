package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=reactjs&paymentid=1234"

func main(){
	fmt.Println("Welcome to url")
	fmt.Println(myurl)

	result,err:=url.Parse(myurl)

	checkNilErr(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Query())
	fmt.Println(result.Query().Get("course"))

	for _,val:= range result.Query(){
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=abc",
	}
	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}

func checkNilErr(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		panic(err)
	}
}