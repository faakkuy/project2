package main

import "fmt"

type Gaji struct {
    GajiPokok        int
    PajakPenghasilan float64
    Asuransi         float64
    TotalGaji        int
}

func main() {
    var gaji Gaji
    var pajak,asuransi float64 
    fmt.Scanln(&gaji.GajiPokok,&gaji.PajakPenghasilan,&gaji.Asuransi)
    pajak = (gaji.GajiPokok) * (gaji.PajakPenghasilan /100)
    asuransi = (gaji.GajiPokok) * (gaji.Asuransi /100)
    gaji.TotalGaji = gaji.GajiPokok - pajak - asuransi
    fmt.Print(gaji.TotalGaji)
}


