package main

import "fmt"

// Prosedur untuk menghitung jumlah kemenangan
func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

// Prosedur untuk menghitung jumlah jd
func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

// Prosedur untuk menghitung jumlah kekalahan
func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

// Prosedur untuk menghitung jumlah g, kegolan, dan selisih g
func hitungGolKegolSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg += (g - k)
}

// Prosedur untuk menghitung jumlah poin
func hitungPoin(jm, jd int, jp *int) {
	*jp = (jm * 3) + jd
}

func main() {
	var N, g, k, jm, jd, jk, jgol, jkegol, jsg, jpoint int

	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Scan(&g, &k)

		// Hitung statistik pertandingan
		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jk)
		hitungGolKegolSelisih(g, k, &jgol, &jkegol, &jsg)
	}

	// Hitung jumlah poin
	hitungPoin(jm, jd, &jpoint)

	// Tampilkan hasil rekap pertandingan
	fmt.Println(N,jm,jd,jk,jgol,jkegol,jsg,jpoint)
}
