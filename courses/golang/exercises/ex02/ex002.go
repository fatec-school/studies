package main

import "fmt"

func main() {
	numbers := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	firstSum := 0
	secondSum := 0

	for _, num := range numbers {
		if num <= 5 {
			firstSum += num
		}else if num >= 6 {
			secondSum += num
		}
	}

	fmt.Println(firstSum)
	fmt.Println(secondSum)
}
