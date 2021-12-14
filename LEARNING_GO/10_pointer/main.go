package main

import "fmt"

//pointers
func main() {
	a := 5
	var b *int = &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//use star to dereference a
	//same in c++
	fmt.Println(*b)
	*b = 23
	fmt.Println(a)
}
