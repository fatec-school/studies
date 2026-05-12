package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	_, err := os.Open("file.txt")
	if err != nil {
		panic("file not found")
	}

}

func main() {
	ReadFile()

	fmt.Println("END")
}
