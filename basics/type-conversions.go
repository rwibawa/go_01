package main

import (
	"math"
	"fmt"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*x))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
