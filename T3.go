package main

import (
	"fmt"
	"math"
)

func panjangX(R,T float64)float64{
	var radian float64
	radian = T * (math.Pi / 180.0)
	return R * math.Cos(radian)

}

func panjangY(R,T float64)float64{
	var  radian float64
	radian = T * (math.Pi / 180.0)
	return R * math.Sin(radian)
}

func main(){
	var R,T,hasilX,hasilY float64
	fmt.Scan(&R,&T)

	hasilX = panjangX(R, T)
	hasilY = panjangY(R, T)

	fmt.Printf("%.2f\n", hasilX)
	fmt.Printf("%.2f\n", hasilY)

}