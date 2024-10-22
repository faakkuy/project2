package main

import "fmt"

func isiArray(n int) [100]int {
	/* mengembalikan array yang berisi bilangan bulat */
	var arr [100]int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr[i] = num
	}
	return arr
}

func countPositiveSumNegative(arr [100]int) (int, int) {
	/* mengembalikan jumlah bilangan positif dan negatif dari array */
	positiveCount := 0
	negativeSum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positiveCount++
		} else if arr[i] < 0 {
			negativeSum += arr[i]
		}
	}
	return positiveCount, negativeSum
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := isiArray(n)
	positiveCount, negativeSum := countPositiveSumNegative(arr)
	fmt.Println(positiveCount, negativeSum)
}
