package main

import (
	"fmt"
)

const nMax int = 100

type dataMobil [nMax]string

func isiData(arrMobil *dataMobil) {
	/*I.S. Terdefinisi nilai array dataMobil yang masih kosong
	F.S. arrMobil berisi data yang diinputkan*/
	var warna string
	index := 0
	for {
		fmt.Scanln(&warna)
		if warna == "-1" {
			break
		}
		(*arrMobil)[index] = warna
		index++
	}
}

func hitung(arrMobil dataMobil) string {
	/*Mengembalikan string warna mobil tertentu yang paling banyak
	melintasi jalan*/
	var n, nMerah, nHitam, nAbu int
	var terbanyak string

	n = len(arrMobil)
	for i := 0; i < n; i++ {
		if arrMobil[i] == "merah" {
			nMerah++
		} else if arrMobil[i] == "hitam" {
			nHitam++
		} else if arrMobil[i] == "abu" {
			nAbu++
		}
	}

	if nMerah >= nHitam && nMerah >= nAbu {
		terbanyak = "merah"
	} else if nHitam >= nMerah && nHitam >= nAbu {
		terbanyak = "hitam"
	} else {
		terbanyak = "abu"
	}
	return terbanyak
}


func main() {
	var arrMobil dataMobil
	isiData(&arrMobil)
	warnaTerbanyak := hitung(arrMobil)
	fmt.Println("Mobil terbanyak :", warnaTerbanyak)
}
