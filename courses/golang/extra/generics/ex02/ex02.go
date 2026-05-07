package main

// Exercício 2 — Função Contains
/*
Crie uma função genérica que verifica se um slice contém um valor.

Assinatura:

func Contains[T comparable](items []T, value T) bool

*/

import (
	"fmt"
	"slices"
)

func Contains[T comparable](items []T, value T) bool {
	return slices.Contains(items, value)
}

func main() {
	numbers := []int{1, 3, 4, 5, 6}
	fmt.Println(Contains(numbers, 1))
}
