package studentdetails

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Marks float32
}

func PrintStudent() map[int]Student {
	students := map[int]Student{
		1: {"ravi", 25, 54.23},
		2: {"sanji", 23, 50.05},
		3: {"zoro", 21, 23.12},
	}
	return students

}

func PrintFullName() {
	// using bufio and os stdin
	fmt.Println("we are importing this file from studentdetails/student.go to print full name using bufio, os Stdin")
	fmt.Println("enter ur full name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	
	fmt.Printf("Hello, %s", name)
}
