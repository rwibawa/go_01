package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
func Sqrt(x float64) float64 {
	z := float64(1)
	for ;z < 10; z++ {
		z := z - (math.Pow(z, 2) - x) / (2*z)
		fmt.Println(z)
	}

	return z
}
