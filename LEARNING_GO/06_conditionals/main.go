package main

import (
	"fmt"
	"time"
)

func main() {
	x := 15
	y := 10

	//if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "green"

	if color == "red" {
		fmt.Printf("color is %v\n", color)
	} else if color == "blue" {
		fmt.Printf("color is not red but %v\n", color)
	} else {
		fmt.Printf("color is not red or blue, it is %v\n", color)
	}

	//switch statements
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today is Wed")
	case today + 1:
		fmt.Println("Wenesday is tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Far away")
	}
}
