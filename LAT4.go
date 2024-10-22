package main

import "fmt"

func main(){
	var n, m int
	fmt.Scan(&n, &m)
	cetak(n,m)
}

func cetak(n, m int){
	if n<m {
		for i := n; i <= m ; i++  {
			fmt.Print(i, " ")
		}
	} else {
		for i := m; i <= n ; i++ {
			fmt.Print(i, " ")
		}
	}
}