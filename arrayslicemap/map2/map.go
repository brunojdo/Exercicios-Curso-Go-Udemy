package main

import (
	"fmt"
)

func main() {
	funcsESalarios := map[string]float64{
		"José":        1133.25,
		"Ana Maria":   2555.00,
		"Pedro Silva": 1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
