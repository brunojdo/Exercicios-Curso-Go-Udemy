package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Galvao Silva":   152200.88,
			"Geraldo Santos": 52223.55,
		},

		"J": {
			"Jo Soares": 2512.55,
			"Jose Lima": 4152.99,
		},
		"P": {
			"Pedro Júnior": 2512.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
