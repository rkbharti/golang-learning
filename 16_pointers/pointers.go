package main

import "fmt"

// changenum demonstrates PASS BY VALUE
// It receives a copy of the variable, so original value will NOT change
func changenum(num int) {
	num = 2
	fmt.Println("[changenum] Inside function (copy changed):", num)
}

// double demonstrates PASS BY REFERENCE (pointer)
// It modifies the original variable using its memory address
func double(num *int) {
	*num = *num * 2
	fmt.Println("[double] Inside function (original updated):", *num)
}

func main() {
	num := 5

	fmt.Println("Initial value in main:", num)

	// 🔹 PASS BY VALUE
	fmt.Println("\n--- Calling changenum (Pass by Value) ---")
	changenum(num)

	fmt.Println("[main] After changenum (original unchanged):", num)

	// 🔹 PASS BY REFERENCE
	fmt.Println("\n--- Calling double (Pass by Reference) ---")
	double(&num)

	fmt.Println("[main] After double (original changed):", num)
}