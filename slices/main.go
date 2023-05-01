package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Orange", "Grapes", "Mango", "Banana"}
	fmt.Printf("Fruit list: %T\n", fruitList);

	fruitList[0] = "Pineapple"
	fmt.Println("Fruit list: ", fruitList);
	fruitList = append(fruitList, "Kiwi");
	fmt.Println("Fruit list: ", fruitList);
	fruitList = append(fruitList[:3]);
	fmt.Println("Fruit list: ", fruitList);


	higScores := make([]int,4)
	higScores[0] = 100
	higScores[1] = 200
	higScores[2] = 300
	higScores[3] = 400
	higScores = append(higScores, 700,500,200,800)
	fmt.Println("High scores: ", higScores);
	sort.Ints(higScores)
	fmt.Println("High scores: ", higScores);

	sort.Slice(higScores, func(i, j int) bool {
		return higScores[i] > higScores[j]
	});

	fmt.Println("High scores: ", higScores);

	var courses = []string{"C#", "Java", "Python", "Go", "JavaScript"}
	fmt.Println("Courses: ", courses);
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses: ", courses);
}