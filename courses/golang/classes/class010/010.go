package main

import "fmt"

func main() {
	var isCar bool = true
	var carValue float64 = 1000.00

	if isCar {
		carValue += 55.50
	} else {
		carValue += 20.50
	}

	fmt.Printf("Final value is %.2f", carValue)
}
