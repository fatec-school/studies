package main

import "fmt"


func main() {
	// Variables not using pointers
	x := 5
	y := x

	y = 10
	fmt.Println(x, y)
	fmt.Println(&x, &y)

	
	// Variables using pointers
	a := 5
	b := &a

	*b = 10
	fmt.Println(a, *b)
	fmt.Println(&a, &b)
}