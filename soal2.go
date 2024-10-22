package main

import (
	"fmt"
)

func countNumber(a, b int) int {
	if keduanyagenap(a, b) {
		return hasilKali(a, b)
	} else if keduanyaganjil(a, b) {
		return hasilJumlah(a, b)
	} else {
		return 0
	}
}

func keduanyaganjil(a, b int) bool {
	return cekBothGanjil(a) && cekBothGanjil(b)
}

func keduanyagenap(a, b int) bool {
	return cekBothGenap(a) && cekBothGenap(b)
}

func cekBothGanjil(a int) bool {
	return a%2 == 1
}

func cekBothGenap(a int) bool {
	return a%2 == 0
}

func hasilKali(a, b int) int {
	return a * b
}

func hasilJumlah(a, b int) int {
	return a + b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(countNumber(a, b))

}