package main

import (
	"fmt"
	"sync"
)

// 4. Contador seguro

// Crie uma variável global:

// var counter int

// Lance 100 goroutines.
// Cada uma incrementa counter 1000 vezes.

// Sem Mutex, o valor final provavelmente ficará errado.

// Depois:

// proteja a variável usando sync.Mutex
// o resultado final deve ser 100000

var counter int

func main() {
	var (
		mx sync.Mutex
		wg sync.WaitGroup
	)

	for range 100 {

		wg.Go(func() {
			for range 1000 {
				mx.Lock()
				counter++
				mx.Unlock()
			}
		})
		
	}

	wg.Wait()

	fmt.Println(counter)
}
