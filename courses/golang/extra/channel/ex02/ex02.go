package main

import (
	"fmt"
	"time"
)

// 2. Producer / Consumer

// Implemente:

// um producer que gera strings
// um consumer que lê do channel

// Exemplo de saída:

// Produzido: tarefa-1
// Consumido: tarefa-1

// Objetivos:

// separar responsabilidades
// entender sincronização natural dos channels

func main() {
	channel := make(chan string)

	go producer(channel)
	go consumer(channel)

	time.Sleep(2 * time.Second)
}

func producer(channel chan string) {
	for i := 1; i <= 10; i++ {
		task := fmt.Sprintf("tarefa-%d", i)

		fmt.Println("Produzido:", task)

		channel <- task
	}

	close(channel)
}

func consumer(channel chan string) {
	for value := range channel {
		fmt.Println("Consumido:", value)
	}
}
