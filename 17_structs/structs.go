// struct is custom data type(group of related field)

package main

import "fmt"

//make struct
type student struct {
	fname string
	lname string
	age   int
	marks float32
}

func printStudent(s student){
	fmt.Println("FirstName: " , s.fname)
	fmt.Println("LastName",s.lname)
	fmt.Println("Age is : ",s.age)
	fmt.Println("Marks is : ",s.marks)

}

func main() {
	//create object of struct
	student1 := student{
		fname: "sameer",
		lname: "raaz",
		age:   25,
		marks: 89.3,
	}
	student1.marks=90
	// fmt.Println(student1)
	student2 := student{
		fname: "ravi",
		lname: "kumar",
		age:   23,
		marks: 55.55,
	}

	 printStudent(student1)
	 fmt.Println("*********************")
	printStudent(student2)
}
