package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{11, 22, 33, 44, 55} // compilador conta e define a quantidade de elementos

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
