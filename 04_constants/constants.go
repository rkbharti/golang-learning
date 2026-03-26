package main

import "fmt"

func main() {
	// const name = "Ravi"
	// const age = 18
	// const isAdult = true
	// const float= 0.5
	// // wee can't change the const name onece we declare 
	// // name = "kumar" 
	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(isAdult)
	// fmt.Println(float)

	// grouping constant by using const()
	const(
		name = "ravi"
		age = 25
		isPass = true
		marks= 89.5
	)

	fmt.Println(name, age , isPass, marks)

}