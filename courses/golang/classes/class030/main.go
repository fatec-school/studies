package main

import (
	"aula30/model"
	"fmt"
)

func main() {
	fmt.Println("Starting...")

	address := model.Address{
		Street: "Main Street",
		Number: 123,
		City:   "Anytown",
	}

	fmt.Printf("Address: %+v\n", address)
}
