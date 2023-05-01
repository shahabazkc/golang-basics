package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)
	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"
	languages["de"] = "German"
	languages["it"] = "Italian"
	languages["ru"] = "Russian"
	languages["cn"] = "Chinese"
	fmt.Println("Languages: ", languages["en"])
	delete(languages, "ru")
	fmt.Println("Languages: ", languages)

	// loops through the map
	for key,value:=range languages{
		fmt.Printf("Key: %v, Value: %v \n", key, value)
	}
}