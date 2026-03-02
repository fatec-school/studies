package main

import "fmt"

func main() {
	for numBase := 1; numBase <= 10; numBase++ {
		for i := range 11 {
			fmt.Printf("%d x %d = %d\n", numBase, i, numBase * i)
		}
	}
}
