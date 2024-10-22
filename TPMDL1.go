package main

import "fmt"

func a(r float64, l, k *float64) {
	*l = 3.14 * r * r
	*k = 2 * 3.14 * r
}

func b(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungtotal(r, s float64, LL, LP, KL, KP, toLuas, totKel *float64) {
	a(r, LL, KL)
	b(s, LP, KP)
	*toLuas = *LL + *LP
	*totKel = *KL + *KP
}

func main() {
	var radius, sisi, LL, LP, KL, KP, toLuas, totKel float64

	var allResults [][]float64

	for {
		fmt.Scan(&radius, &sisi)

		if radius == 0 && sisi == 0 {
			break
		}

		hitungtotal(radius, sisi, &LL, &LP, &KL, &KP, &toLuas, &totKel)

		result := []float64{radius, sisi, LL, LP, KL, KP, toLuas, totKel}
		allResults = append(allResults, result)
	}

	if len(allResults) > 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")

		for _, result := range allResults {
			fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", result[0], result[1], result[2], result[3], result[4], result[5], result[6], result[7])
		}
	}
}