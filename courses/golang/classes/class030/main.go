package main

import (
	"aula30/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting...")

	address := model.Address{
		Street: "Main Street",
		Number: 123,
		City:   "Anytown",
	}

	person := model.Person{
		Name:    "John Doe",
		Address: address,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}

	fmt.Printf("Person: %+v\n", person)
	//fmt.Printf("Age: %d", model.GetAge(person))
	fmt.Printf("Age: %d\n", person.GetAge())

	fmt.Printf("Name: %s\n", person.Name)
	person.ChangeName("Mick")
	fmt.Printf("New name: %s\n", person.Name)
}
