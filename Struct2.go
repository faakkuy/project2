package main

import "fmt"

// Definisikan tipe data Balok dengan komponen panjang, lebar, tinggi, dan ukuran
type Balok struct {
    Panjang int
    Lebar   int
    Tinggi  int
    Ukuran  Geometri
}

// Definisikan tipe data Geometri untuk menyimpan nilai luas alas dan volume
type Geometri struct {
    LuasAlas int
    Volume   int
}

func main() {
    // Input ukuran panjang, lebar, dan tinggi balok
    var balok Balok
    fmt.Scan(&balok.Panjang)
    fmt.Scan(&balok.Lebar)
    fmt.Scan(&balok.Tinggi)

    // Hitung luas alas dan volume balok
    balok.Ukuran.LuasAlas = balok.Panjang * balok.Lebar
    balok.Ukuran.Volume = balok.Panjang * balok.Lebar * balok.Tinggi

    // Tampilkan hasil
    fmt.Println(balok.Ukuran.LuasAlas, balok.Ukuran.Volume)
}