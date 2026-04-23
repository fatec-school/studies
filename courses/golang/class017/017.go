package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var lowNumbers = make([]int, 0) 
	var lowNumbers []int

	/*
	for i := 0; i < len(numbers); i++ {
			if numbers[i] < 5 {
				lowNumbers = append(lowNumbers, numbers[i])
			}
	}
	*/

	for i, valor := range numbers {
		if valor < 5 {
			lowNumbers = append(lowNumbers, valor)
			fmt.Printf("Position [%d] %v\n", i, lowNumbers)
		}
	}

}
 