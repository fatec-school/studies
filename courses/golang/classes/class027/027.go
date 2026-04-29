package main

import "fmt"

var globalPseudoNumber = 42

func main() {
	a, b := 10, 5
	sum, diff, prod, quot := operations(a, b)

	fmt.Println("Sum:", sum-globalPseudoNumber)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot)
}

func operations(a, b int) (sum int, diff int, prod int, quot int) {
	pseudoNumber := 42

	return a + b + pseudoNumber, a - b, a * b, a / b
}
