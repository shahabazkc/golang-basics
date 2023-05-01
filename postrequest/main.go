package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func main(){
	fmt.Println("Welcome to post Request");
	PerformPostRequest("http://localhost:3000/post")
	performPostJsonRequest("http://localhost:3000/postform")
}


func PerformPostRequest(url string){
	fmt.Println("PerformPostRequest")

	requestBody := strings.NewReader(`{"name":"abc","age":20}`)
	fmt.Println(requestBody)
	response,err := http.Post(url, "application/json", requestBody)

	checkNilErr(err)

	defer response.Body.Close()
  	content,err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println("Content: ", string(content))

}

func performPostJsonRequest(urlValue string){
	data := url.Values{}
	data.Add("name", "abc")
	data.Add("age", "20")
	response,err := http.PostForm(urlValue,data)
	checkNilErr(err)
	defer response.Body.Close()

	content,err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("Content: ", string(content))
}

func checkNilErr(err error){
	if err != nil{
		fmt.Println("Error: ", err)
		panic(err)
	}
}