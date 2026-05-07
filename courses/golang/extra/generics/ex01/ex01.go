package main

import "fmt"

// Exercício 1 — Função Max
/* Crie uma função genérica que receba dois valores e retorne o maior deles.

### Requisitos

- Deve funcionar para:
  - int
  - float64
  - string
- Use constraints.

*/

type numbers interface {
	int | float64 | string
}

func Grower[T numbers](x T, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	a := 2.1
	b := 4.2

	fmt.Println(Grower(a, b))
}
