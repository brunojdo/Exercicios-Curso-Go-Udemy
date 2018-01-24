package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}

	y := []int{}
	z := x[:0]

	fmt.Println(x)

	//z := []int{}

	y = append(y, 100, 200)
	z = append(z, 44)

	fmt.Println(y, x, z)

	/* 	z = append(z, 1000, 2000)
	   	x[1] = 11

	   	fmt.Println(x)

	   	w := x[:0]

	   	x[1] = 22

	   	w = append(w, 13, 12)

	   	fmt.Println(w)
	   	//fmt.Println(cap(x), cap(y), cap(z)) */
}
