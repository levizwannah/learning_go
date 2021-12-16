package main

import (
	"fmt"
	"strconv"
)

// structs have two types of methods
// pointer receivers and value receivers

//defining a struct
type Person struct {
	firstName string
	lastName  string
	gender    string
	age       int
}

//methods are outside the struct

//Greeting method value receiver
func (p Person) greeting() string {
	return "Hello " + p.firstName + " " + p.lastName
}

//Greeting introduction
func (this Person) introduce() string {
	return "Hello my name is " + this.firstName + " " + this.lastName + ". I am " + strconv.Itoa(this.age) + " years old."
}

//pointer receivers
func (this *Person) birthday() {
	this.age++
}

//pointer receiver
func (this *Person) getMarried(spouseLastName string) {
	if this.gender == "m" {
		return
	}

	this.lastName = spouseLastName
}

func main() {
	//initializing a struct
	person2 := Person{firstName: "Levi", lastName: "Zwannah", gender: "m", age: 21}
	person3 := Person{firstName: "Joline", lastName: "Pinky", gender: "f", age: 20}

	//equally
	// person2 := Person{"Levi", "Zwannah", "M", 21}
	//fmt.Println(person2)
	fmt.Println(person2.greeting())
	fmt.Println(person2.introduce())
	person2.birthday()
	person3.getMarried(person2.lastName)
	fmt.Println(person3.introduce())
}
