package main

import "fmt"

func main() {
	fmt.Println(add(42,53))
}
func add(x int, y int) int {
	return x + y
}