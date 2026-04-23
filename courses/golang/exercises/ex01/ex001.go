package main

import "fmt"


func main() {
	var numbers = [2]int{2,3}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	fmt.Println(sum)
}