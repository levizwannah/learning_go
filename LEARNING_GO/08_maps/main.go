package main

import "fmt"

func main() {
	//maps
	emails := make(map[string]string)

	emails["levi"] = "levi@mail.com"
	emails["sharon"] = "sharon@mail.com"
	emails["john"] = "john@gmail.com"

	delete(emails, "levi")
	fmt.Println(emails)

	//create map and add key values
	ages := map[string]int{"levi": 10, "mary": 20}
	fmt.Println(ages)
	ages["mike"] = 67
	fmt.Println(ages)
}
