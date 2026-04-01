// // maps is an assoscited data struture and helpe isn fast lookup dor data instead of using loop

// package main

// import (
// 	"fmt"
// 	"maps"
// )
// func main(){
// 	fmt.Println("using map to store data in  key and values and we use this when we dont know how many elemts are store here  ")
// 	m:= make(map[string]string)
// 	m["fname"] = "GO" //fname is key , "GO" is values
// 	m["lname"] = "LANG"
// 	fmt.Println(m["fname"], m["lname"])
// 	fmt.Println("length of map is : ",len(m))
// 	fmt.Println("for this type of value whic is not store it give empty element in string case []",m["phone"]) //if key not stored then it give ZEROTH VALUE

// 	// now we check what zeroth element given wehne we have int and string as key -values
// 	m1:=make(map[string]int)
// 	m1["age"] = 20
// 	m1["age2"] = 30
// 	fmt.Println("because of  existence key it gave : ",m1["age"]) //20
// 	fmt.Println("because of non existence key it gave: ",m1["phone"]) //0 because of non existence key, so zeroth value in int

// 	// delete map
// 	delete(m,"lname")
// 	delete(m1,"age2")
// 	fmt.Println("here it will show nondeleted key from m: ",m , "here it will show  nondeleted key from m2",m1 )

// 	// create map without make function , when element is already given
// 	fmt.Println("create map without make function , when element is already given ")
// 	m2:=map[string]bool{"isAdult":true, "isTeenager" : false}
// 	fmt.Println("it will print m2 element with key and value both : ",m2)

// 	// check element in map to take some action

// 	m3:=map[string]int{"phone": 20, "price": 2000}
// 	k, v := m3["price"]
// 	fmt.Println("print the key value of what we have written like price: 2000 : ",k)
// 	if v {
// 		fmt.Println("All good")
// 	}else{
// 		fmt.Println("NOT ok!")
// 	}

// 	// check equal in map
// 	m4:=map[string]int{"phone": 2, "price": 2000}
// 	m5 :=map[string]int{"phone": 2, "price": 200}
// 	fmt.Println(maps.Equal(m4 , m5))

// }

package main

import "fmt"

func main() {
	students := map[int]string{
		1: "Ravi",
		2: "Anil",
	}
	for {
		var choice int
		fmt.Println("******************************************")
		fmt.Println("enter your choice")
		fmt.Println("1: Add student")
		fmt.Println("2: view all student")
		fmt.Println("3: search by ID")
		fmt.Println("4: Delete student by id ")
		fmt.Println("5: Exit")
		fmt.Print("YOU TYPE : ")
		fmt.Scanln( &choice)
		fmt.Println("******************************************")

		switch choice {
		case 1:
			fmt.Println("Add New Student")
			var id int
			var name string
			for {
				fmt.Println("Enter the id to Add New Student")
				fmt.Scanln(&id)
				_, exists := students[id]
				if exists {
					fmt.Println("id already exist")
				} else {
					break
				}
			}
			fmt.Println("enter new Student name ")
			fmt.Scanln(&name)
			students[id] = name
			fmt.Println("Added Succesfully ")
			fmt.Println("******************************************")

		case 2:
			fmt.Println("View All Student")
			for key, value := range students {
				fmt.Printf("ID: %d | Name: %s\n", key, value)
			}
			fmt.Println("******************************************")
		case 3:
			fmt.Println("Search by ID")
			var id int
			fmt.Println("Enter the id to Search ")
			fmt.Println("---")
			fmt.Println("  |")
			fmt.Println("  |")
			fmt.Println("  v")
			fmt.Scanln(&id)
			name, exist := students[id]
			if exist{
				fmt.Println(id, "details is " , name)
			}else{
				fmt.Println("NO such ID FOUND ...")
			}
			fmt.Println("******************************************")
		case 4:
			fmt.Println("Delete")
			var id int
			fmt.Println("Enteer the ID to delete ")
			fmt.Scanln(&id)
			_,exist := students[id]
			if exist {
				delete(students, id)
				fmt.Println("Deleted Succesfullly", id)
			} else{
				fmt.Println("NO such ID found")
			}
			fmt.Println("******************************************")
		case 5:
			fmt.Println("EXisting ....")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}

}
