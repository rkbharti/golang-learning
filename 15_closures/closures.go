package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the number")
	fmt.Scanln(&num)

	increment:= func ()int  {
		num += 10
		return num
	}

fmt.Println(increment())
fmt.Println(increment())
fmt.Println(increment())

}