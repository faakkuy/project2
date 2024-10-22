package main

import "fmt"

type Pemain struct {
    Nama   string
    Gol    int
    Assist int
}

func main() {
    var pemain [3]Pemain
    for i := 0; i < 3; i++ {
        fmt.Scan(&pemain[i].Nama, &pemain[i].Gol, &pemain[i].Assist)
    }

    golTerbanyak := pemain[0]
    for i := 1; i < 3; i++ {
        if pemain[i].Gol > golTerbanyak.Gol {
            golTerbanyak = pemain[i]
        }
    }

    assistTerbanyak := pemain[0]
    for i := 1; i < 3; i++ {
        if pemain[i].Assist > assistTerbanyak.Assist {
            assistTerbanyak = pemain[i]
        }
    }

    fmt.Println(golTerbanyak.Nama)
    fmt.Println(assistTerbanyak.Nama)
}