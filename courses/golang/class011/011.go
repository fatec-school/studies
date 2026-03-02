package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	for i := range 10 {
		fmt.Println(i)
	}

	word := "word"
	for i := 0; i < len(word); i++ {

		if string(word[i]) == "r" {
			break
		}

		if string(word[i]) == "o" {
			continue
		}

		fmt.Printf("%s ", string(word[i]))
	}
}
