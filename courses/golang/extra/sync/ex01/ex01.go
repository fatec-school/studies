package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Esperar múltiplas goroutines

// Crie 5 goroutines.
// Cada uma deve:

// imprimir "Trabalhando X"
// dormir entre 1 e 3 segundos
// imprimir "Finalizado X"

// Use Go sync.WaitGroup para garantir que o main() espere todas terminarem.

func main() {
	var wg sync.WaitGroup

	for range 5 {

		wg.Go(func() {

			fmt.Println("Trabalhando X")
			time.Sleep(time.Second)
			fmt.Println("Finalizado X")
		})
	}

	wg.Wait()
}
