package main

import "fmt"

func main() {
	var t1 float64 = 2.5
	var t2 float64 = 3.2
	var alo float64
	amo := t2 + t1
	fmt.Printf("result %f", amo)

	var number int = 12

	for i := 0; i < number; i++ {

		alo += amo + t1
	}
	fmt.Printf("result2 : %f", alo)
}
