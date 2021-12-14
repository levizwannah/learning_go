package main

import "fmt"

//use to loop through arrays, maps, and slices
func main() {

	ids := []int{10, 20, 30, 40, 50, 45}

	for index, id := range ids {
		fmt.Printf("%d - ID: %d\n", index, id)
	}

	//not using the index put an underscore

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//suming
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	//range with maps
	ages := map[string]int{"levi": 10, "mary": 40, "hame": 45}

	for key, value := range ages {
		fmt.Printf("%s age is %d\n", key, value)
	}

	//get the keys only
	for key := range ages {
		fmt.Println("Key: ", key)
	}
}
