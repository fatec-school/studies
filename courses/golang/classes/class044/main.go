package main

import "fmt"

func main() {
	channel := make(chan int)
	go setList(channel)

	for value := range channel {
		fmt.Println(value)
	}
}

func setList(channel chan int) {
	for num := range 100 {
		channel <- num
	}

	defer close(channel)
}
