package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{}
var wg sync.WaitGroup // pointer
var mut sync.Mutex // pointer
func main() {
	// go greeter("Hello")
	// greeter("world")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _,web := range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done()
	result,err := http.Get(endpoint)

	if err !=nil{
		log.Fatal(err)
	} else {
		mut.Lock()
		fmt.Printf("%d status code for website for %s \n",result.StatusCode,endpoint)
		signals = append(signals, endpoint)
		mut.Unlock()
	}
}