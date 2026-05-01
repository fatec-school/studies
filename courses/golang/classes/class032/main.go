package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func ShowArea(g geometry) {
	fmt.Println(g.area())
}

func main() {
	fmt.Println("starting...")

	rectangle := rectangle{
		width:  1,
		height: 2,
	}

	circle := circle{
		radius: 3,
	}

	ShowArea(rectangle)
	ShowArea(circle)
}
