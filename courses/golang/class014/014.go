package main

import "fmt"

func main() {
	list := []int{4, 9, 6, 7}

	fmt.Printf("Full list: %v\n", list)
	fmt.Printf("First element: %v\n", list[0])
	fmt.Printf("Last element: %v\n", list[len(list)-1])

	fmt.Printf("List length: %v\n", len(list))
}
