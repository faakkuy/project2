package main

import "fmt"

type pegawai struct {
	nama      string
	gajiPokok int
	masaKerja float64
	besarBonus int
}

func main() {
	var p pegawai

	// Masukan data pegawai
	fmt.Scan(&p.nama,&p.gajiPokok,&p.masaKerja)

	// Hitung bonus
	hitungBonus(&p)

	// Cetak besar bonus
	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %d\n", p.nama, p.besarBonus)
}

func hitungBonus(p *pegawai) {
	if p.masaKerja >= 10 {
		p.besarBonus = p.gajiPokok * 3 / 2
	} else if p.masaKerja >= 5 {
		p.besarBonus = p.gajiPokok * 3 / 4
	} else {
		p.besarBonus = p.gajiPokok / 2
	}
}
