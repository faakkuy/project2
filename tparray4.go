package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)
	fmt.Println(getFibonaci(x))
}

func getFibonaci(x int) int {
	var d1, d2, temp, t, i int
	if x == 1 {
		return 0
	}
	if x == 2 {
		return 1
	}
	d1 = 1
	d2 = 1
	temp = 0
	t = 1
	for i = 1; i <= x-2; i++ {
		t = t + d2
		temp = d2
		d2 = d2 + d1
		d1 = temp
	}
	return t
}