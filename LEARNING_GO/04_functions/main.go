package main

import "fmt"

//function without params
func hello() string {
	return "Hello, World"
}

//function with param
func helloParam(name string) string {
	return "Hello " + name
}

//function without return
func helloVoid() {
	fmt.Println("Hello, Void")
}

func helloMultiParams(num1, num2 int) int {
	return num1 + num2
}

func helloMultiRetur() (string, int) {
	return "Hello World", 5
}

func main() {
	fmt.Println(hello(), helloParam("Levi"), helloMultiParams(23, 45))
	helloVoid()
}
