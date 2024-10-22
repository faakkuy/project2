package main

import "fmt"

func reamur(c float64) float64 {
	r := c * 0.80
	return r
}

func fahrenheit(c float64) float64 { 
	f := (9 * c) / 5 + 32
	return f
}

func kelvin(c float64) float64 {
	k := c + 273
	return k
}

func main() {
	var ca, ck, s float64
	
	fmt.Scan(&ca, &ck, &s)
	fmt.Printf("%10s %10s %10s %10s\n", "celcius", "Reamur", "Fahrenheit", "Kelvin") 
	
	for ca <= ck {
		c := ca
		r := reamur(ca)
		f := fahrenheit(ca)
		k := kelvin(ca)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", c, r, f, k)
		ca = ca + s
	}
}