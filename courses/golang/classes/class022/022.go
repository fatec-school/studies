package main

import "fmt"

func main() {
	maps := map[string]int{"sp": 99999, "rj": 10000, "pn": 120}

	fmt.Println(maps)
	delete(maps, "pn")
	fmt.Println(maps)
}
