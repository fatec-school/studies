package main

import "fmt"

// 1. Channel básico

// Crie:

// uma goroutine que envia números de 1 até 5
// a main recebe e imprime os valores

// Objetivos:

// usar make(chan int)
// enviar (<-)
// receber (<-ch)

func main() {
	channel := make(chan int)
	go setValues(channel)

	for value := range channel {
		fmt.Println(value)
	}
}

func setValues(channel chan int) {
	for num := range 5 {
		channel <- num
	}
	defer close(channel)
}
