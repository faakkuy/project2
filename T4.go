package main

import "fmt"

func main(){
	var hasil int
	hasil = 0
	for i := 0; i <= 50; i++ {
		hasil = hasil + ((1/(i + 1)) - (1/(i + 2)))
	}
	fmt.Println(hasil)
}