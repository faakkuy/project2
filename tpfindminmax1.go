package main

import "fmt"

// Deklarsi konstanta 20 elemen
const NMAX int = 20

// Deklarasi tabInt dengan total NMAX elemen
type tabInt [100]int

func main() {
    // Deklarasi variabel bertipe tabInt
	var arr tabInt
	var n int
	// Deklarasi banyak elemen array
	// Pembacaan data bilangan
	baca(&arr, &n)
	// Cetak elemen array
	cetakElemen(arr, n)
	// Cetak Info
	cetakInfo(arr, n)
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca bilangan bulat dan memasukkan bilangan bulat positif 
		        ke dalam array A.Pembacaan dilakukan selama terbaca 
		        bilangan positif dan n < NMAX.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	for i := 0; i < 100; i++ {
		fmt.Scan(&A[i])
		*n += 1
		if A[i] <= 0 || i == 20{
			break
		}
	}
}

func cetakElemen(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Elemen array: <e1> <e2> <e3> ... <en>"
	*/
	fmt.Print("Elemen array: ")
	for i := 0; i < n-1; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen maksimum dari array A dengan banyak elemen n */
	var max int
	max = A[0]
	for i := 0; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	
	return max
}

func minimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen minimum dari array A dengan banyak elemen n */
	var min int
	min = A[0]
	if A[19] > 0 {
		for i := 0; i < n; i++ {
			if A[i] < min {
				min = A[i]
			}
		}
	} else {
		for i := 0; i < n-1; i++ {
			if A[i] < min {
				min = A[i]
			}
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Nilai maksimum, nilai minimum, dan banyaknya elemen tercetak dengan format:
			"Nilai maksimum: <max_value>
			 Nilai minimum: <min_value>
			 Banyak elemen: <n>"
	*/
	fmt.Printf("Nilai maksimum: %d\n", maksimum(A, n))
	fmt.Printf("Nilai minimum: %d\n", minimum(A, n))
	if A[19] > 0 {
		n = 20
		fmt.Printf("Banyak elemen: %d", n)
	} else {
		fmt.Printf("Banyak elemen: %d", n-1)
	}
}