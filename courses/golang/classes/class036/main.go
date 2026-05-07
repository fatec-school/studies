package main

import "fmt"

func main() {
	numbers := []int{2, 3, 5, 6, 6}
	words := []string{"a", "b", "c"}

	fmt.Println(reverse(numbers))
	fmt.Println(reverse(words))
}

type customConstraint interface {
	int | string
}

func reverse[T customConstraint](slice []T) []T {
	newInts := make([]T, len(slice))
	newIntsLen := len(slice) - 1

	for i := range slice {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newInts
}
