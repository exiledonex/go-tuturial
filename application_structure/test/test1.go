package main

import "fmt"

func main() {
	var A int32 = 12
	var b float32 = 7.5
	var c int32 = 100
	minos := float32(A)*b + b - float32(c)
	fmt.Printf("the power of a & b is %f \n", minos)
}
