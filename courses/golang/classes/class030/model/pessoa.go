package model

import "time"

type Person struct {
	Name      string
	BirthDate time.Time
	Address   Address
}

/*
func GetAge(p Person) int {
	birthYear := p.BirthDate.Year()
	actualYear := time.Now().Year()

	return actualYear - birthYear
}
*/

func (p Person) GetAge() int {
	birthYear := p.BirthDate.Year()
	actualYear := time.Now().Year()

	return actualYear - birthYear
} 