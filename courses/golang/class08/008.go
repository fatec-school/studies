package main

import "fmt"

func main() {
	salary := 900.00
	var salaryWithBonus float64

	salaryWithBonus = salary

	fmt.Printf("Salary: %f\n", salaryWithBonus)

	if salary < 1000 {
		salaryWithBonus += 100
		fmt.Printf("Salary less 1000. New salary with bonus is %.2f\n", salaryWithBonus)
	}
}
