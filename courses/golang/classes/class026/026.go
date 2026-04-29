package main

import "fmt"

func main() {
	sum, diff, prod, quot := operation(10, 5)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot)
}

func operation(num1 int, num2 int) (sum int, diff int, prod int, quot int) {
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2
}
