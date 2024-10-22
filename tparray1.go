package main

import (
	"fmt"
)

const NMAX int = 10

type Mahasiswa struct {
	Nama    string
	NIM     int
	Jurusan string
}

type MahasiswaArr [NMAX]Mahasiswa

// Function to find the Mahasiswa with the maximum Jurusan lexicographically
func findMax(A MahasiswaArr, n int) int {
	maxIndex := 0
	for i := 0; i < n; i++ {
		if A[i].Jurusan > A[maxIndex].Jurusan {
			maxIndex = i
		}
	}
	return maxIndex
}

// Function to find a Mahasiswa by NIM using sequential search
func findSequentialNIM(A MahasiswaArr, n, x int) string {
	for i := 0; i < n; i++ {
		if A[i].NIM == x {
			return A[i].Nama
		}
	}
	return "-1"
}

// Function to find a Mahasiswa by NIM using binary search
func findBinary(A MahasiswaArr, n, x int) string {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if A[mid].NIM == x {
			return A[mid].Nama
		} else if A[mid].NIM < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return "-1"
}

func main() {
	var A MahasiswaArr
	// Initialize the Mahasiswa array (A) with some data
	A = MahasiswaArr{
		{Nama: "Alice", NIM: 101, Jurusan: "Computer Science"},
		{Nama: "Bob", NIM: 102, Jurusan: "Mathematics"},
		{Nama: "Charlie", NIM: 103, Jurusan: "Physics"},
		// Add more Mahasiswa as needed
	}

	n := 3 // Number of Mahasiswa entries

	// Example usage of findMax
	maxIndex := findMax(A, n)
	fmt.Printf("Mahasiswa with the lexicographically largest Jurusan: %s\n", A[maxIndex].Nama)

	// Example usage of findSequentialNIM
	nimToFind := 102
	name := findSequentialNIM(A, n, nimToFind)
	if name != "-1" {
		fmt.Printf("Mahasiswa with NIM %d: %s\n", nimToFind, name)
	} else {
		fmt.Printf("Mahasiswa with NIM %d not found\n", nimToFind)
	}

	// Example usage of findBinary
	// Note: The array should be sorted by NIM for binary search to work correctly
	name = findBinary(A, n, nimToFind)
	if name != "-1" {
		fmt.Printf("Mahasiswa with NIM %d: %s\n", nimToFind, name)
	} else {
		fmt.Printf("Mahasiswa with NIM %d not found\n", nimToFind)
	}
}
