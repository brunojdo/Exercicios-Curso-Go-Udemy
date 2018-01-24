package main

import (
	"fmt"
)

/* func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}

} */

func notaParaConceitoFatorada(n float64) string {
	switch n {
	case 9, 10:
		return "A"
	case 8:
		return "B"
	case 5, 6, 7:
		return "C"
	case 3, 4:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceitoFatorada(5))
	fmt.Println(notaParaConceitoFatorada(9))
	fmt.Println(notaParaConceitoFatorada(1))
}
