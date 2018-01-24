package main

import (
	"fmt"
)

func main() {
	a1 := [...]int{1, 2, 3, 4, 5}
	a2 := [3]int{11, 22, 33}

	fmt.Println(a1)
	fmt.Println(a2)

	s1 := a1[0:2]
	fmt.Println(s1)

	a1[0] = 40

	fmt.Println(s1)

}
