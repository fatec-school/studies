package main

import "fmt"

func main() {
	numbers := []int{1, 12, 4, 5, 7, 4, 6}

	firstNumbers := numbers[:3] // 3 first items
	lastNumbers := numbers[4:] // After 4 item
	lastItem := numbers[len(numbers) - 1:] // Last Item

	fmt.Println(firstNumbers)
	fmt.Println(lastNumbers)
	fmt.Println(lastItem)
}
