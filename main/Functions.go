package main

import "fmt"

func Functions() {
	a, b := multiReturns(5, "Hello")
	fmt.Println(a, b)

	c, d := swap(5, 10)
	fmt.Println(c, d)

	fmt.Println(recursive(5))

	dfs(3, "")
	fmt.Println(arr)
}

func swap(a int, b int) (x int, y int) {
	x = b
	y = a
	return
}

func multiReturns(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

var arr []string

func dfs(n int, word string) {
	if n == 0 {
		arr = append(arr, word)
		return
	}
	for i := 0; i < 4; i++ {
		var next = word + string(rune('a'+i))
		dfs(n-1, next)
	}
}
