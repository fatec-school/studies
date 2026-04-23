package main

import "fmt"

func main() {
	maps := map[string]int{"sp": 90000, "rj": 10990}

	for index, value := range maps {
		fmt.Printf("%s -> %d\n", index, value)
	}
}
