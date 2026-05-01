package model

import (
	"time"
)

type ShoppingList struct {
	BuyDate time.Time
	Market  string
	Items   []Item
}

func (sp *ShoppingList) Init(market string, date time.Time) *ShoppingList {
	sp.BuyDate = date
	sp.Market = market

	var items []Item
	items = append(items, Item{
		Name:  "Item 1",
		Price: 10,
	})

	items = append(items, Item{
		Name:  "Item 2",
		Price: 20,
	})

	sp.Items = items

	return sp
}
