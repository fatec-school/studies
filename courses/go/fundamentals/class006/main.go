package main

import "fmt"

func main() {
	var name string = "Nickolas"
	fmt.Println(name)

	var anotherName = "Another name"
	fmt.Println(anotherName)

	var (
		variable1 string = "variable 1"
		variable2 string = "variable 2"
	)
	variable3, variable4 := "variable 3", "variable 4"

	fmt.Println(variable1, variable2, variable3, variable4)

	const constant string = "const"
	fmt.Println(constant)
}
