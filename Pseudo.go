package main

import "fmt"

// tipe bentukan struct waktu dengan field jam, menit, detik semuanya bertipe integer
type waktu struct {
	jam, menit, detik int
}

// tipe bentukan struct kendaraan dengan field jenis kendaraan (jke), nopol (no),
// jam masuk (jm) dan jam keluar (jk) bertipe waktu
// bea parkir (bea) dan durasi parkir (durasi) bertipe integer
type kendaraan struct {
	jke               string
	no                string
	jm, jk            waktu
	bea, durasi       int
}

func main() {
	var k01, k02, k03 kendaraan
	var jenis, nopol string
	var j1, m1, d1, j2, m2, d2 int

	fmt.Scan(&jenis, &nopol, &j1, &m1, &d1, &j2, &m2, &d2)
	// Membuat kendaraan k01
	k01 = kendaraan_baru(jenis, nopol, j1, m1, d1, j2, m2, d2)

	fmt.Scan(&jenis, &nopol, &j1, &m1, &d1, &j2, &m2, &d2)
	// Membuat kendaraan k02
	k02 = kendaraan_baru(jenis, nopol, j1, m1, d1, j2, m2, d2)

	fmt.Scan(&jenis, &nopol, &j1, &m1, &d1, &j2, &m2, &d2)
	// Membuat kendaraan k03
	k03 = kendaraan_baru(jenis, nopol, j1, m1, d1, j2, m2, d2)

	// hitung bea parkir 3 kendaraan
	bea_parkir(&k01)
	bea_parkir(&k02)
	bea_parkir(&k03)

	// cetak data parkir per kendaraan
	cetak_data(k01)
	cetak_data(k02)
	cetak_data(k03)
}

func kendaraan_baru(j, n string, a, b, c, d, e, f int) kendaraan {
	/*  Mengembalikan nilai kendaraan dengan data jenis kendaraan (j), nopol (n),
	jam masuk dalam jam (a), menit (b), detik (c) dan jam keluar dalam jam (d), menit (e), detik (f).
	*/
	var k kendaraan
	k.jke = j
	k.no = n
	k.jm = waktu{a, b, c}
	k.jk = waktu{d, e, f}
	return k
}

func durasi_parkir(k *kendaraan) {
	/*  IS: k terdefinisi, kecuali durasi parkir (k.durasi) dan bea parkir (k.bea)
	    Proses : Ubah durasi tot_detik menjadi k.durasi dalam satuan jam
	    FS: Durasi parkir (k.durasi) berisi nilai dalam satuan jam
	*/
	var tot_detik int
	tot_detik = (k.jk.jam*3600 + k.jk.menit*60 + k.jk.detik) - (k.jm.jam*3600 + k.jm.menit*60 + k.jm.detik)
	k.durasi = tot_detik / 3600
}

func bea_parkir(k *kendaraan) {
	/*  IS: k terdefinisi, kecuali durasi parkir (k.durasi) dan bea parkir (k.bea)
	Proses: Memanggil prosedur durasi_parkir dan menghitung k.bea
		    Jika kendaraan mobil, biaya parkir jam pertama Rp 5000, per jam selanjutnya Rp 2500
			Jika kendaraan motor, biaya parkir jam pertama Rp 2000, per jam selanjutnya Rp 1000
	FS: Durasi parkir (k.durasi) berisi nilai dalam satuan jam
	*/
	durasi_parkir(k)
	if k.jke == "mobil" {
		k.bea = 5000 + (k.durasi-1)*2500
	} else if k.jke == "motor" {
		k.bea = 2000 + (k.durasi-1)*1000
	}
}

func cetak_data(k kendaraan) {
	/*
		IS: k terdefinisi
		FS: Tercetak data parkir kendaraan dengan format:
			"Bea parkir <jke> nopol <no> dengan durasi <durasi> jam: Rp <bea>."
	*/
	fmt.Printf("Bea parkir %s nopol %s dengan durasi %d jam: Rp %d.\n", k.jke, k.no, k.durasi, k.bea)
}
