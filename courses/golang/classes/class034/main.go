package main

import "fmt"

func main() {
	fmt.Println("starting...")

	var a interface{} // Is a any type
	var b any // Other any type

	a = 1
	b = "asdf"

	fmt.Println(a)
	fmt.Println(b)

	var slice []any
	slice = append(slice, 12)
	slice = append(slice, "oi")

	fmt.Println(slice)
}
