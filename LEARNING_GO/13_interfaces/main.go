package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	height, width float64
}

func (this Circle) area() float64 {
	return this.radius * this.radius * math.Pi
}

func (this Rectangle) area() float64 {
	return this.height * this.width
}

func getArea(this Shape) float64 {
	return this.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 15}
	rec := Rectangle{width: 20, height: 10}

	fmt.Printf("The area of circle is %f\n", getArea(circle))
	fmt.Printf("The area of the rectangle is %f\n", getArea(rec))
}
