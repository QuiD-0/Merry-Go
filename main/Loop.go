package main

import "fmt"

func Loop() {

	for i := 0; i < 5; i += 2 {
		fmt.Println("i =", i)
	}
	println("======")

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i =", i)
	}
	println("=======")

	for i := 0; i < 5; i++ {
		if i%2 == 1 {
			break
		}
		fmt.Println("i =", i)
	}
	println("=======")

	list := []string{"a", "b", "c", "d"}
	for idx, val := range list {
		fmt.Println(idx, val)
	}

}
