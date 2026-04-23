package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
	fmt.Printf("Full list: %v\n", list)

	list = append(list, 5)
	fmt.Printf("New full list: %v\n", list)

	stringList := []string{"a", "b", "c"}
	fmt.Printf("Full String list: %v\n", stringList)

	intList := make([]int, 1)
	fmt.Printf("%T\n", intList)

	for i := 0; i < len(list); i++ {
		fmt.Printf("%v,", list[i])
	}
}
