package main

import (
	"ex03/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting...")

	shoppingList := model.ShoppingList{
		BuyDate: time.Date(2026, 1, 1, 1, 0, 0, 0, time.Local),
		Market:  "Market XPTO",
		Items: []model.Item{
			model.Item{
				Name:  "Item 1",
				Price: 99,
			},
			model.Item{
				Name:  "Item 2",
				Price: 123,
			},
		},
	}

	fmt.Printf("Shopping List: %+v", shoppingList)
}
