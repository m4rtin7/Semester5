package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 0.0
	acc := 0.00000000001

	for x-z*z > acc {
		z += x / 1000
	}
	return z
}

func main() {
	fmt.Println(Sqrt(100))
}
