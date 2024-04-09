package main

import "fmt"

func Output() {
	a, b, c := 1, 2, 3

	//stdout

	fmt.Print("Hello, World!\n")
	fmt.Println("Hello, World!")
	fmt.Println(a, "-", b, "-", c)
	fmt.Println("Hello, World!")

	//stderr

	println("Hello, World!")
	println(a, "-", b, "-", c)
}
