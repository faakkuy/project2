package main

import "fmt"

func main(){
	var hujanTurun, punyaPayung string
	fmt.Scan(&hujanTurun)
	fmt.Scan(&punyaPayung)
    goNoGo(hujanTurun, punyaPayung)
}

func goNoGo(hujanTurun string, bawaPayung string) {
	if hujanTurun == "ya" {
		if bawaPayung == "ya"{
			fmt.Println("berangkat")
		}else{
			fmt.Println("diam di rumah")
		}
	}else{
		fmt.Println("berangkat")
	}

}