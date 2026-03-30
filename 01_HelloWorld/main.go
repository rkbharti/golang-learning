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

	fmt.Println("********************************************************************************")

	fmt.Println("using bufio and os.Stdin to print a full statemnt till user got to next line")
	studentdetails.PrintFullName()
}
