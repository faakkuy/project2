package main

import (
	"fmt"
)

type Penjumlahan struct {
	jumKonsonan int
	jumVokal    int
}

const NMax int = 1000

var deretKonsonan [NMax]byte
var deretVokal [NMax]byte

func isKonsonan(c byte) bool {
	return (c >= 'a' && c <= 'z') && !isVokal(c)
}

func isVokal(c byte) bool {
	vokal := "aiueoAIUEO"
	for i := 0; i < len(vokal); i++ {
		if c == vokal[i] {
			return true
		}
	}
	return false
}

func main() {
	var status = true
	var penjumlahan Penjumlahan
	var idxKonsonan, idxVokal int

	for status {
		var inputKarakter byte
		fmt.Scanf("%c", &inputKarakter)

		if inputKarakter != '\n' && inputKarakter != ' ' {
			if inputKarakter == '.' {
				status = false
			} else if isKonsonan(inputKarakter) {
				deretKonsonan[idxKonsonan] = inputKarakter
				idxKonsonan++
				penjumlahan.jumKonsonan++
			} else if isVokal(inputKarakter) {
				deretVokal[idxVokal] = inputKarakter
				idxVokal++
				penjumlahan.jumVokal++
			}
		}
	}

	fmt.Printf("%s\n", string(deretKonsonan[:idxKonsonan]))
	fmt.Println(penjumlahan.jumKonsonan)
	fmt.Printf("%s\n", string(deretVokal[:idxVokal]))
	fmt.Println(penjumlahan.jumVokal)
}
