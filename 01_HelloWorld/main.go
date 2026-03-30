package main

import (
	"fmt"
	"golang_tutotorial/studentdetails"
)

func main() {
	fmt.Println("Hello! Ravi")
	students := studentdetails.PrintStudent()
	for _, s := range students {
		fmt.Println("PRRINT FROM OTHER FILES USING STRUCT TO SHOW STUDENT ", s.Name, s.Age, s.Marks)
	}
}
