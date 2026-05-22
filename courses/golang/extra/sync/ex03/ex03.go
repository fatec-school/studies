package main

import (
	"fmt"
	"sync"
	"time"
)

// 3. Download simulado

// Simule downloads concorrentes:

// files := []string{"foto.png", "video.mp4", "musica.mp3"}

// Cada goroutine deve:

// imprimir "Baixando X"
// esperar tempo aleatório
// imprimir "X concluído"

// Use WaitGroup.

func main() {
	var wg sync.WaitGroup

	files := []string{"foto.png", "video.mp4", "musica.mp3"}

	wg.Add(len(files))

	for _, file := range files {
		func() {
			go fakeDownload(file, &wg)
		}()
	}

	wg.Wait()
}

func fakeDownload(file string, wg *sync.WaitGroup) {
	fmt.Printf("Baixando %s\n", file)

	time.Sleep(3 * time.Second)

	fmt.Printf("%s concluido\n", file)

	wg.Done()
}
