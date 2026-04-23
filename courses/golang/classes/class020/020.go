package main

import "fmt"

func main() {
	maps := map[string]int{"sp": 10, "cg": 90}
	var emptyMap = map[string]int{}

	emptyMap["Maya"] = 19

	fmt.Println(maps)
	fmt.Println(emptyMap)

	/*
		value, exist := maps["rj"]
		if exist {
			fmt.Println(value)
		}
	*/

	if value, exist := maps["rj"]; exist {
		fmt.Println(value)
	}
	
}
