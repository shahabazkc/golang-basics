package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request")

	PerformGetRequest("http://localhost:3000")
}

func PerformGetRequest(url string) {
	fmt.Println("Performing get request")
	response,err:=http.Get(url)

	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode);
	fmt.Println("Content length: ", response.ContentLength);

	var responseString  strings.Builder
	content,err:=ioutil.ReadAll(response.Body)
	checkNilErr(err)

	byteCount,err:=responseString.Write(content)
	checkNilErr(err)


	fmt.Println("Byte count: ", byteCount)
	fmt.Println("Content: ", responseString.String())

	checkNilErr(err)
	fmt.Println("Content: ", string(content))

}

func checkNilErr(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		panic(err)
	}
}