package main

import (
	"fmt"
)

// hitungJumlah merupakan prosedur untuk menghitung jumlah bilangan bulat dari 1 hingga N
func hitungJumlah(N int, S *int) {
	// Inisialisasi nilai awal S
	*S = 0

	// Menghitung jumlah bilangan dari 1 hingga N
	for i := 1; i <= N; i++ {
		*S += i
	}
}

func main() {
	var N, S int

	// Input nilai N dari pengguna
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&N)

	// Memanggil prosedur hitungJumlah untuk menghitung jumlah dari 1 hingga N
	hitungJumlah(N, &S)

	// Menampilkan hasil penjumlahan
	fmt.Println(S)
}
