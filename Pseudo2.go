package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    hasil := sukuN(N)
    fmt.Println(hasil)
}

func sukuN(N int) int {
    if N <= 3 {
        return 1
    }
    return sukuN(N-1) + sukuN(N-2) + sukuN(N-3)
}
