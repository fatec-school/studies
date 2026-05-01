package main

import "time"

type shoppingList struct {
	BuyDate time.Time
	Market  string
	Items   []item
}
