package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, t := 1.0, 0.0
	for i := 0; i < 10; i++ {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-8 {
			break
		}
	}
	return z
}

func main() {
	fmt.Printf("Sqrt: %v\n", Sqrt(2))
	fmt.Printf("math.Sqrt: %v\n", math.Sqrt(2))
}
