package main

import (
	"fmt"
	"sync"
)

// 2. Processamento paralelo de números

// Dado um slice:

// nums := []int{1,2,3,4,5,6,7,8,9,10}

// Crie uma goroutine para cada número que:

// calcule o quadrado
// imprima o resultado

// Espere todas finalizarem usando WaitGroup.

func main() {
	var wg sync.WaitGroup
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(len(nums))

	for _, num := range nums {
		func() {
			go exponentiation(num, &wg)
		}()
	}

	wg.Wait()
}

func exponentiation(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)

	wg.Done()
}
