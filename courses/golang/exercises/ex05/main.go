package main

import (
	"ex03/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting...")

	var shoppingList model.ShoppingList
	shoppingList.Init("Mercado legal", time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local))

	fmt.Printf("Shopping List: %+v", shoppingList)
}
