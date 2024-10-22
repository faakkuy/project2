package main

import "fmt"

func main() {
	var angkaDitebak int
	fmt.Scan(&angkaDitebak)

	tebakAngka(angkaDitebak)
}

func tebakAngka(angkaDitebak int) {
	var tebakan int
	attempt := 0
	guessCorrect := false

	for !guessCorrect {
		fmt.Scan(&tebakan)
		attempt++

		if tebakan == angkaDitebak {
			guessCorrect = true
		} else if tebakan/10 == angkaDitebak/10 || tebakan%10 == angkaDitebak%10 || tebakan == angkaDitebak/10 || tebakan == angkaDitebak%10 {
			guessCorrect = true
		}
	}

	fmt.Println( attempt)
}
