package main

func main() {
	a, b := 10, 5
	sum, difference, product, quotient := operations(a, b)

	println("Sum:", sum)
	println("Difference:", difference)
	println("Product:", product)
	println("Quotient:", quotient)
}

func operations(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
