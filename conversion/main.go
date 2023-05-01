package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Weclome to conversion")
	fmt.Println("Please rate our service")

	reader:= bufio.NewReader(os.Stdin)
	input, _:= reader.ReadString('\n')
	fmt.Println("Your rating is: ",input)


	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64);
	if err != nil {
		fmt.Println("Error: ",err)
	}else {
		numRating++
	}
	fmt.Println("Your rating is: ",numRating)
	
}