package main

import "fmt"

//channels are a typed pipe through which you can send and receive values
//with the channel operator <-
//buffered channels ch := make(chan int, 2)

//sum dividing the work between two go routines
func sum(s []int, c chan int) {
	sum := 0
	for _, value := range s {
		sum += value
	}

	c <- sum //send sum to c
}

func main() {
	c := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	//get the result and sum
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
