package main

import (
	"fmt"
)

func main() {
	// simple switch where we assign a value
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("no is : ", (i), " not in anyone of the case")
	// }

	// multiple condition switch
	// switch time.Now().Weekday(){
	// case time.Saturday , time.Sunday :
	// 	fmt.Println("hurrah!it's weekend")
	// default :
	// 	fmt.Println("oops! it weekkkkk-day")
	// }

	// type switch where it help to find the type of variable by making a function callable 

	whoAmI := func(i interface{}){
		switch i.(type){
		case int:
			fmt.Println(i , "is an integer")
		case string:
			fmt.Println(i,"is a string")
		case bool:
			fmt.Println(i, "is a boolena")
		default:
			fmt.Println(i,"maybe float or something check twice")
		}
	}
	whoAmI(0.5)
}