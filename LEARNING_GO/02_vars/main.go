package main

import "fmt"

//Variables
//string
//bool
//int
//int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64
//byte - same as uint8
//rune - same as int32
//float32 float64
//complex64 complex128

func main() {
	//vars declaration with var
	var name string = "Levi"
	var name2 = "Mary"
	const THE_CONST = true

	fmt.Println(name + " " + name2)

	var age = 45
	fmt.Printf("%T\n", age)

	//shorthand declaration. This can only be used in a function
	shortHand := 120
	fmt.Println(shortHand)

	//multiple variables declaration
	model, make := "Tuzine", 567
	fmt.Println(model, make)
}
