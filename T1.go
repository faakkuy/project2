package main

import "fmt"

// Procedure untuk menghitung banyaknya kemunculan angka genap di kedua dadu
func countEvenOccurrences(dice1, dice2 int, evenCount *int) {
    if dice1%2 == 0 && dice2%2 == 0 {
        *evenCount++
    }
}

func main() {
    var dice1, dice2, evenCount int

    // Membaca input sampai muncul angka sama-sama ganjil
    for {
        fmt.Scanln(&dice1,&dice2)

        // Memanggil prosedur untuk menghitung kemunculan angka genap
        countEvenOccurrences(dice1, dice2, &evenCount)

        // Periksa apakah muncul angka sama-sama ganjil, jika ya, keluar dari loop
        if dice1%2 != 0 && dice2%2 != 0 {
            break
        }
    }

    // Menampilkan jumlah kemunculan angka genap
    fmt.Println(evenCount)
}
