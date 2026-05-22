package main

import (
	"fmt"
	"sync"
)

// 5. Banco simples

// Crie uma struct:

// type BankAccount struct {
//     balance int
// }

// Implemente:

// Deposit(amount int)
// Withdraw(amount int)

// Use Mutex para evitar race conditions.

// Teste com várias goroutines depositando e sacando simultaneamente.

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.balance < amount {
		fmt.Println("sem saldo")
		return
	}

	b.balance -= amount
}

func (b *BankAccount) Balance() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.balance
}

func main() {
	account := BankAccount{
		balance: 1000,
	}

	var wg sync.WaitGroup

	for range 30 {
		wg.Add(2)

		go func() {
			defer wg.Done()
			account.Deposit(300)
		}()

		go func() {
			defer wg.Done()
			account.Withdraw(300)
		}()
	}

	wg.Wait()

	fmt.Println("Saldo final:", account.Balance())
}