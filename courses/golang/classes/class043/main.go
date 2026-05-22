package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	
	var m sync.Mutex
	i := 0

	for x := range 10000 {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()

		x++
	}

	time.Sleep(5 * time.Second)

	fmt.Println(i)
}

