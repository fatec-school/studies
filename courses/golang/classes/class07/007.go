package main

import(
	"strconv"
	"fmt"
)

func main() {
	// Basic type conversion

	// var number8 int8
	// var number int
	// number = int(number8)
	// print(number)

	// String to int conversion
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		println("Error converting string to int:", err)
	} else {
		println("Converted number:", num)
		fmt.Printf("Value type: %T\n", num)
	}

	strBool := "true"
	boolVal, err := strconv.ParseBool(strBool)
	if err != nil {
		println("Error converting string to bool:", err)
	} else {
		println("Converted bool:", boolVal)
		fmt.Printf("Value type: %T\n", boolVal)
	}

	strFloat := "4.2"
	floatValue, err := strconv.ParseFloat(strFloat, 1)
	if err != nil {
		println("Error converting string to float: ", err)
	} else {
		println("Converted float: ", floatValue)
		fmt.Printf("Value type: %T\n", floatValue)
	}
}
