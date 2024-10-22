package main

import "fmt"

func fungsiY(x float64)float64{
    // fungsi akan mengembalikan nilai y = 1/x jika 
    var hasil float64
    hasil = 1/x
    return hasil
}



func main (){
    var x float64
    fmt.Scan(&x)
    fmt.Println(fungsiY(x))
}