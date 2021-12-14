package main

import "fmt"

func main() {

	//long method
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short method
	for i := 0; i <= 10; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	//fizzbuzz
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
