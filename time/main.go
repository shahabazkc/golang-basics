package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time calculation")
	today:= time.Now()
	fmt.Println("Start time: ",today.Format(("02-01-2006 03:04:05 Monday")))

	createdDate:= time.Date(2020,time.August, 0, 0, 0, 0, 0, time.UTC);
	fmt.Println("Created date: ",createdDate.Format("02-01-2006 03:04:05 Monday"));

	fmt.Println("Started time: ",time.Now().Format("02-01-2006 03:04:05 Monday"));
	time.Sleep(5 * time.Second);
	fmt.Println("Ended time: ",time.Now().Format("02-01-2006 03:04:05 Monday"));
}
