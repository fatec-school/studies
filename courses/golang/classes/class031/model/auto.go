package model

type Auto struct {
	Year  int
	Model string
}

type Motorcycle struct {
	Auto             // Extends in golang]
	CylinderCapacity int
}

type Car struct {
	Auto // Extends in golang
	Door int
}
