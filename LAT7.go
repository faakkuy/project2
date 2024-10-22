package main

import "fmt"

func main(){
    var x,y int
	fmt.Scan(&x,&y)
	(swipe(x,y))
}
func swipe(ref1,ref2 int){
	fmt.Scan(&ref1,&ref2)
	for ref1 != ref2 {
		if ref1 > ref2 {
			fmt.Println(ref2,ref1)
		}else{
			fmt.Println(ref1,ref2)
		}
	}
}