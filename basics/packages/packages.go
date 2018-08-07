package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// same number
	fmt.Println("My favorite number is", rand.Intn(10))

	// setting seed
	rand.Seed(10)
	fmt.Println("My favorite number is", rand.Intn(10))
}
