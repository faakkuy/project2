package main

import "fmt"

// Define the function to calculate the total
func lingkaran(r int) int {
	var total int
	total = 22 * r / 7
	return total
}

func main() {
	// Declare the radius variable
	var r int

	// Assign a value to r
	r = 5

	// Call the function and store the result
	total := lingkaran(r)

	// Print the result
	fmt.Println("Total:", total)
}
