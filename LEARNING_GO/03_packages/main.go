package main

import (
	"fmt"
	"github.com/levizwannah/LEARNING_GO/03_packages/strutil"
	"math"
)

func main() {
	fmt.Println(math.Abs(-233), math.Sqrt(49), math.Floor(89.9), math.Ceil(78.65))
	fmt.Println(strutil.Reverse("EATING EVERY THING"))
}
