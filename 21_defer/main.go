package main

import "fmt"

func main() {
	// now we differ the value by using difer
	fmt.Println("---------------------")
	defer fmt.Println("---------------------")
	defer fmt.Println("Hello") //Run this line at the end of the function”
	defer fmt.Println("World") //Run this line at the end of the function”
	
}
