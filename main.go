package main

import "fmt"

func add(a, b int) int {
	fmt.Println("rising cov")
	fmt.Println("rising cov")
	fmt.Println("rising cov")
	fmt.Println("rising cov")
	return a + b
}

func main() {
	fmt.Println(add(1, 1))
	fmt.Println("downing cov")
	fmt.Println("downing cov")
}
