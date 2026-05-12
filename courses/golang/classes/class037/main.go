package main

import "os"

func main() {
	_, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
}
