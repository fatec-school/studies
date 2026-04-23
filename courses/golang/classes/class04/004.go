package main

import (
	"fmt"
	"reflect"
)

func main() {
	num1 := 10
	num2 := 20

	result := num1 + num2
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

	// fmt.Println(num1 + num2)
}
