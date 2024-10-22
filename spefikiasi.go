package main

import "fmt"

func main() {
	var jawaban string
	fmt.Println(confess(jawaban))

}

func confess(jawaban string) string {
	fmt.Println("aku sebenarnya suka kamu apakah kamu mau jadi pacarku ????")
	fmt.Scan(&jawaban)
	if jawaban == "yes" {
		return ("letsgoooo")
	} else {
		return ("okeyyyy")
	}
}
