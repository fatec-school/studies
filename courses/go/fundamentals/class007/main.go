package main

import (
	"errors"
	"fmt"
)

func main() {
	var num01 int8 = 127
	fmt.Println(num01)

	var num02 uint8 = 255
	fmt.Println(num02)

	var num03 float32 = 10.7
	fmt.Println(num03)

	var str string = "text"
	fmt.Println(str)

	var boolean bool = true
	fmt.Println(boolean)

	var error error = errors.New("My internal error")
	fmt.Println(error)
}
