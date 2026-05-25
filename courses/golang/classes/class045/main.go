package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 100) // O 100 indica o tamanho do channel
	go setList(channel)

	for value := range channel {
		fmt.Println("recebendo: ", value)
		time.Sleep(time.Second)
	}
}

// func setList(channel <-chan int) { o <-chan indica que o channel é só de leitura
// 	for i := range 100 {
// 		channel <- i
// 	}
// 	close(channel)
// }

// func setList(channel chan <- int) { o chan<- indica que o channel é só de escrita
// 	for i := range 100 {
// 		channel <- i
// 	}
// 	close(channel)
// }

func setList(channel chan int) {
	for i := range 100 {
		channel <- i
		fmt.Println("enviando: ", i)
	}
	close(channel)
}
