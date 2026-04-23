package main

import "fmt"

func main() {
	list := []int{4, 9, 6, 7}

	fmt.Println("List: ", list)
	fmt.Println("List in 1 position: ", list[0])
	fmt.Println("List in 2 position: ", list[1])
	fmt.Println("List length: ", len(list))

	list = append(list, 10)
	fmt.Println("New list: ", list)



	stringList := []string{"Go", "Java", "PHP"}
	stringList = append(stringList, "Rust")
	fmt.Println("String list: ", stringList)



	otherList := make([]int, 1)
	otherList = append(otherList, 1)
	otherList = append(otherList, 2)
	fmt.Println("Other list: ", otherList)
	
	sum := 0
	for i := 0; i < len(otherList); i++ {
		fmt.Println(otherList[i])
		sum += otherList[i]
	}

	fmt.Printf("Avg other list: %d\n", sum / len(otherList))

}
