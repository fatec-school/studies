package main

import (
	"fmt"
)

func main() {
	word := "word"
	length := len(word)
	
	i := 0
	for i < length {
		fmt.Printf("%s ", string(word[i]))
		i++
	}
}
