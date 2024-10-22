package main

import "fmt"

type Vehicle struct {
    Jenis   string 
    Durasi  int  
}

func main() {
	var jenisKendaraan string
	var durasiParkir int
	fmt.Scanln(&jenisKendaraan, &durasiParkir)
    totalPendapatan := 0

    for (jenisKendaraan == "motor") && (jenisKendaraan == "mobil" )|| (durasiParkir > 0) {
		fmt.Scanln(&jenisKendaraan, &durasiParkir)
       
		var tarifPerJam int
        if jenisKendaraan == "motor" {
            tarifPerJam = 1000
        } else {
            tarifPerJam = 5000
        }
        jumlahBayar := tarifPerJam * durasiParkir

        totalPendapatan += jumlahBayar
    }

    fmt.Println( totalPendapatan)
}
