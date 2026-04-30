package main

import (
	"aula31/model"
	"fmt"
)

func main() {
	fmt.Println("starting...")

	autoMoto := model.Auto{
		Year: 2022,
		Model: "XPTO",
	}

	motorcycle := model.Motorcycle{
		Auto: autoMoto,
		CylinderCapacity: 129,
	}

	fmt.Printf("Auto: %+v\n", motorcycle)
	fmt.Printf("Year: %d\n", motorcycle.Year)
	fmt.Printf("Model: %s", motorcycle.Model)
}