package main

import (
	"fmt"
)

func main() {
	fmt.Print("mesma ")
	fmt.Println("linha")
	fmt.Println("nova linha")

	x := 3.141516

	//fmt.Println("O valor de x é")

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

}
