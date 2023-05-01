package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")
	// for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i is: ", i)
	// }
	// while loop
	j := 0
	for j < 10 {
		j++
		if j == 3 {
			goto middle
		}
		if j == 5 {
			continue
		}
		fmt.Println("j is: ", j)
	}
	middle:
		fmt.Println("Middle")
	// do while loop
	// k := 0
	// for {
	// 	fmt.Println("k is: ", k)
	// 	k++
	// 	if k == 5 {
	// 		break
	// 	}
	// }

	var days = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Days: ", days)
	// for i:=0;i<len(days);i++{
	// 	fmt.Println("Day: ", days[i])
	// }
	// for i := range days{
	// 	fmt.Println("Day: ", i)
	// }

	for index, day := range days {
		fmt.Printf("Index: %v, Day: %v \n", index, day)
	}

	


}
