package main

import "fmt"

func init() {
	// Is like a constructor, it will be executed before the main function
	fmt.Println("Init function")
}

func main() {
	fmt.Println("Main function")
}