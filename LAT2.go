package main

import "fmt"

func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrimeZ(bilangan))
}

func isPrimeZ (bilangan int)bool{
	var hasil int
	if bilangan < 1 {
		return false
	}
	hasil = 0
	for i := 1; i <= bilangan; i++{
		if bilangan % i == 0{
			hasil = hasil + 1
		}
	}
	return hasil % 2 != 0 
}