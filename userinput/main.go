package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")

	text, _ := reader.ReadString('\n')
	fmt.Println("The text you entered is: ", text)

	//convert string to int
	fmt.Println("Enter a number: ")
	text, _ = reader.ReadString('\n')
	fmt.Println("The number you entered is: ", text)

	num, err := strconv.ParseFloat(strings.TrimSpace(text), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added number: ", num+1)
	}

}
