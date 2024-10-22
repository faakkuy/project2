package main

import "fmt"

type titik struct {
	x, y  float64
	warna string
}

func main() {
	var x1, y1, x2, y2 float64
	var w1, w2 string
	var t1, t2 titik

	// Membaca data x1, y1, dan w1
	fmt.Scanln(&x1,&y1,&w1)

	// Membaca data x2, y2, dan w2
	fmt.Scanln(&x2,&y2,&w2)

	// Membuat titik baru
	t1 = pembuatanTitikBaru(x1, y1, w1)
	t2 = pembuatanTitikBaru(x2, y2, w2)

	// Mencetak data titik
	fmt.Printf("Data titik 1: Koordinat (%v, %v), warna %s\n", t1.x, t1.y, t1.warna)
	fmt.Printf("Data titik 2: Koordinat (%v, %v), warna %s\n", t2.x, t2.y, t2.warna)
}

func pembuatanTitikBaru(x, y float64, w string) titik {
	return titik(x, y, w)
}
