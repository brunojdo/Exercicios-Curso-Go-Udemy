package main

import (
	"fmt"
)

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	div := float64(len(numeros))

	if div != 0 {
		return total / div
	}

	return 0

}

func main() {
	fmt.Printf("Média: %.2f", media(100, 4.44, 7.9, 5, 0, 0))
}
