package main

import "fmt"

func Type() {

	var a bool = true
	var b int = 5
	var c float32 = 3.14
	var d string = "Hi!"

	fmt.Printf("a = %t, b = %d, c = %f, d = %s\n", a, b, c, d)

	var aInt int = 10
	var bInt int8 = 127
	var cInt int16 = 32767
	var dInt int32 = 2147483647
	var eInt int64 = 9223372036854775807

	fmt.Printf("a = %d, b = %d, c = %d, d = %d, e = %d\n", aInt, bInt, cInt, dInt, eInt)
}
