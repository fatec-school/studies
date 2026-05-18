package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := range 1000 {
		go showMessage(strconv.Itoa(i))
	}
}

func showMessage(message string) {
	fmt.Println(message)
}

