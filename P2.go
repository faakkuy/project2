package main

import "fmt"

func secMax(max, secondMax *int) {
	var num int
	fmt.Scan(&num)

	for num != 0 {
		if num > *max {
			*secondMax = *max
			*max = num
		} else if num > *secondMax && num != *max {
			*secondMax = num
		}
		fmt.Scan(&num)
	}
}

func main() {
	// Contoh penggunaan
	var max, secondMax int

	secMax(&max, &secondMax)
	fmt.Println(secondMax)
}
