// student management system using CRUD operations
package main

import "fmt"

// create struct for student detailss
type Student struct{
	ID int
	name string
	age int
	marks float32
}

// 1.0 function to add new student details 
func addStudent(students map[int]Student){
	var id int
	var name string
	var age int
	var marks float32

	
	// check validation id already in map or not
	for {
		fmt.Println("Enter ID: ")
		fmt.Scanln(&id)
		_,exists :=students[id]
		if exists{
			fmt.Println("Id already exist")
		} else {
			break
		}
	}
	fmt.Println("Enter the Student name: ")
	fmt.Scanln(&name)
	
	fmt.Println("Enter the Age: ")
	fmt.Scanln(&age)

	fmt.Println("Enter the Marks: ")
	fmt.Scanln(&marks)

	students[id] = Student{id, name, age, marks}

	fmt.Println("NEW STUDENT DETAILS ADDED")
}


// 2.0  function to print all student details 
func printStudent(s Student){
	fmt.Println("ID is " , s.ID ,"|", "Name is : ", s.name, "|" , "Age is : ", s.age , "|" , "Marks is : ", s.marks)
}

// 3.0 function to update the student details 
func updatedetails(students map[int]Student){
	var id int

		fmt.Println("enter ID to UPDATE : ")
		fmt.Scanln(&id)
		s, exists := students[id]

		if !exists{
			fmt.Println("Student ID not found")
		 return 
		}
		// show Id data
		fmt.Println("student ID  for " , id, "is : " )
		printStudent(s)

		// ask to update using switch
		fmt.Println("\n what u want to edit")
		fmt.Println("1. Name")
		fmt.Println("2. Age")
		fmt.Println("3. Marks")

		var choice int
		fmt.Println("Enter ur choice")
		fmt.Scanln(&choice)
		switch choice{
		case 1:
			var newname string
			fmt.Println("enter the New name ")
			fmt.Scanln(&newname)
			s.name = newname
		case 2:
			var newage int
			fmt.Println("enter the new Age: ")
			fmt.Scanln(&newage)
			s.age = newage
		case 3:
			var newmarks float32
			fmt.Println("enter the updated marks")
			fmt.Scanln(&newmarks)
			s.marks = newmarks
		default:
			fmt.Println("Invalid Choice to Update")

		}
		// save to map 
		students[id]=s
		fmt.Println("Student Data Updated sucessfully")
		
	

}
// 4.0 function to delete the student details 
func deleteStudent(students map[int]Student){
	var id int
	fmt.Println("Enter the Id to delete")
	fmt.Scanln(&id)
	_,exists:= students[id]
	if !exists{
		fmt.Println("No Such ID found ")
		return
	}
	delete(students, id)
	fmt.Println(id," : DELETED SUCCESSFULLY ")

}

func main(){
	var choice int
	// use map to store data and faast lookup and edit and update by key 
	students := map[int]Student{
		1 : {1, "Ravi",23, 56.7},
		2 : {2, "Raaz", 24 , 89.3},
		3 : {3, "Rahul", 25, 69.69},
		4 : {4, "Ritik", 22, 99.9},
	}

	for {
		fmt.Println("\n---MENU----")
		fmt.Println("1. Add Students")
		fmt.Println("2. View All Students")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5 . Exit")
		fmt.Println("enter ur choice : ")
		fmt.Scanln(&choice)

		switch choice{
		case 1: 
			fmt.Println("Add students")
			addStudent(students)
		case 2:
			fmt.Println("View All Students")
			for _,s := range students{
				// fmt.Printf("%T\n", s)
				printStudent(s)
			}
		case 3:
			fmt.Println("update student")
			fmt.Printf("%T\n", students)
			updatedetails(students)
		case 4:
			fmt.Println("Delete Student by UNIQUE ID")
			deleteStudent(students)
		case 5:
			fmt.Println("Existing ...")
			return
		default:
			fmt.Println("YOU type by mistakely : ",choice)

			}
	}


}