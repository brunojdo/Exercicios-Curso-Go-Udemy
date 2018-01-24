package main

import (
	"fmt"
)

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func meuFatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * meuFatorial(n-1)
	}

}

func main() {
	/* resultado, _ := fatorial(5)
	fmt.Println(resultado) */
	fmt.Println(meuFatorial(0))

	/* 	_, err := fatorial(-4)
	   	if err != nil {
	   		fmt.Println(err)
	   	} */

}
