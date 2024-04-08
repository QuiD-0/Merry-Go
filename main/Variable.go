package main

import (
	"fmt"
)

var name = "Variable"

func main() {
	const a int = 10
	var b int = 20
	var c = a + b
	d := 30
	fmt.Printf("a = %d, b = %d, c = %d, d = %d\n", a, b, c, d)

	var isTrue bool = true
	fmt.Printf("isTrue = %t\n", isTrue)

	q, w, e, r := 1, 2, 3, 4
	fmt.Printf("q = %d, w = %d, e = %d, r = %d\n", q, w, e, r)

	var (
		z int
		x int    = 1
		y string = "hello"
	)
	fmt.Printf("z = %d, x = %d, c = %s\n", z, x, y)

	const (
		aa = iota
		bb
		cc
	)
	fmt.Printf("aa = %d, bb = %d cc = %d\n", aa, bb, cc)

}
