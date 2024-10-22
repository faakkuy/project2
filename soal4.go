package main

import "fmt"

func pembagianLoop(n,m int)int{
	var hasil int
	hasil = 0

	for n >= m {
		n = n - m
		hasil++
	}
	return hasil
}


func main () {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(pembagianLoop(n, m))
}