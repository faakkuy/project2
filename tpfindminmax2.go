package main

import "fmt"

const NMAX int = 11

type tabPemain struct {
	nama, nomorPunggung, posisi [100]string
	tinggi                      [100]int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(&timnas, &nPemain)
	// cetak data pemain
	cetakPemain(timnas, nPemain)
	// cetak nama pemain tertinggi
	cetakNamaPemainTertinggi(timnas, nPemain)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
			 	Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	for i := 0; i < 100; i++ {
		fmt.Scan(&A.nama[i])
		if A.nama[i] == "none" || i == 11 {
			break
		}
		fmt.Scan(&A.nomorPunggung[i], &A.posisi[i], &A.tinggi[i])
		*n += 1
	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	fmt.Println("Data Pemain: ")
	if A.nama[10] != "none" {
		for i := 0; i < n; i++ {
			fmt.Printf("%s %s %s %d\n", A.nama[i], A.nomorPunggung[i], A.posisi[i], A.tinggi[i])
		}
	} else {
		for i := 0; i < n-1; i++ {
			fmt.Printf("%s %s %s %d\n", A.nama[i], A.nomorPunggung[i], A.posisi[i], A.tinggi[i])
		}
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var max int
	var cetakMax string
	max = A.tinggi[0]

	for i := 0; i < n; i++ {
		if A.tinggi[i] > max {
			max = A.tinggi[i]
			cetakMax = A.nama[i]
		}
	}

	fmt.Printf("Pemain dengan badan tertinggi: %s\n", cetakMax)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var min int
	var cetakMin string
	min = A.tinggi[0]

	for i := 0; i < n; i++ {
		if A.tinggi[i] < min {
			min = A.tinggi[i]
			cetakMin = A.nama[i]
		}
	}

	fmt.Printf("Pemain dengan badan terendah: %s", cetakMin)
}
