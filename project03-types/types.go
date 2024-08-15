package main

import "fmt"

func main() {
	var c1 = complex64(1 + 2i)
	c2 := complex128(2 + 3i)
	fmt.Println(c1) // (1+2i)
	fmt.Println(c2) // (2+3i)
}
