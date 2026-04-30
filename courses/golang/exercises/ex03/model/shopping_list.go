package model

import "time"

type ShoppingList struct {
	BuyDate time.Time
	Market  string
	Items   []Item
}
