package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing in main file")
	err := checkmail.ValidateFormat("email@gmail.com")
	if err != nil {
		fmt.Println("Email is invalid")
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Println("Email is valid")
}
