package main

import "fmt"

func Array() {

	arr := [...]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := []int{1, 2, 3, 4}
	arr4 := [5]int{1, 2}
	arr5 := [5]int{1: 10, 2: 40}
	arr6 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("arr2 = %v\n", arr2)
	fmt.Printf("arr3 = %v\n", arr3)
	fmt.Printf("arr4 = %v\n", arr4)
	fmt.Printf("arr5 = %v\n", arr5)
	fmt.Printf("arr6 = %v\n", arr6)

	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[2:]

	fmt.Printf("slice = %v\n", slice)
	fmt.Printf("slice2 = %v\n", slice2)

	slice = append(slice, 6)
	fmt.Printf("slice = %v\n", slice)

	slice = append(slice, slice2...)
	fmt.Printf("slice = %v\n", slice)

	fmt.Printf("cap = %d, len = %d\n", cap(slice), len(slice))
	slice = append(slice, 7)
	fmt.Printf("cap = %d, len = %d\n", cap(slice), len(slice))
	slice = append(slice, 8)
	fmt.Printf("cap = %d, len = %d\n", cap(slice), len(slice))

}
