package main

import "fmt"

func main() {
	numbers := []int{2, 3, 5, 6, 6}
	words := []string{"a", "b", "c"}

	fmt.Println(reverse(numbers))
	fmt.Println(reverse(words))
}

func reverse[T int | string](slice []T) []T {
	newInts := make([]T, len(slice))
	newIntsLen := len(slice) - 1

	for i := range slice {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newInts
}
