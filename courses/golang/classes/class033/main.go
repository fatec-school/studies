package main

import (
	"errors"
	"fmt"
)

func ShowError(error error) {
	fmt.Println(error.Error())
}

type NetworkProblem struct {
	rede     bool
	hardware bool
}

func (n NetworkProblem) Error() string {
	if n.rede {
		return "Network error"
	} else if n.hardware {
		return "Hardware problem"
	}
	return "other error"

}

func main() {
	fmt.Println("starting...")

	fmt.Println(errors.New("test"))

	np := NetworkProblem{
		rede:     false,
		hardware: true,
	}

	ShowError(np)
}
