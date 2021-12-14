package main

import "fmt"

func main() {
	var fruitArr [2]string

	fruitArr[0] = "Binana"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr[0], fruitArr[1])

	//delare and assign
	//same for array and slices
	fruits := []string{"Mango", "apple", "pineapple"}
	fmt.Println(fruits, len(fruits))

	//range based
	fmt.Println(fruits[0:2])

}
