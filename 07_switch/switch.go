// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// simple switch where we assign a value
// 	// i := 5
// 	// switch i {
// 	// case 1:
// 	// 	fmt.Println("one")
// 	// case 2:
// 	// 	fmt.Println("two")
// 	// case 3:
// 	// 	fmt.Println("three")
// 	// default:
// 	// 	fmt.Println("no is : ", (i), " not in anyone of the case")
// 	// }

// 	// multiple condition switch
// 	// switch time.Now().Weekday(){
// 	// case time.Saturday , time.Sunday :
// 	// 	fmt.Println("hurrah!it's weekend")
// 	// default :
// 	// 	fmt.Println("oops! it weekkkkk-day")
// 	// }

// 	// type switch where it help to find the type of variable by making a function callable

// 	whoAmI := func(i interface{}){
// 		switch i.(type){
// 		case int:
// 			fmt.Println(i , "is an integer")
// 		case string:
// 			fmt.Println(i,"is a string")
// 		case bool:
// 			fmt.Println(i, "is a boolena")
// 		default:
// 			fmt.Println(i,"maybe float or something check twice")
// 		}
// 	}
// 	whoAmI(0.5)
// }

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var user string
	choices := []string{"rock", "paper", "scissor"}

	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println("Enter your choice: Rock / Paper / Scissor and type 'EXIT' to close the game ")
		fmt.Scanln(&user)
		user = strings.ToLower(user)
		if user == "exit"{
				fmt.Println("GAme OVER..")
				break
		}
		computer := strings.ToLower(choices[rand.Intn(3)])
		fmt.Println("Computer choose", computer)
		

		switch {
			
		
		case user == computer:
			fmt.Println("It's TIE")
		case user == "rock" && computer == "scissor":
			fmt.Println("YOU WON")
		case user == "paper" && computer == "rock":
			fmt.Println("YOU WON")
		case user == "scissor" && computer == "paper":
			fmt.Println("YOU WON")

		case user == "rock" && computer == "paper":
			fmt.Println("COMPUTER WON")
		case user == "paper" && computer == "scissor":
			fmt.Println("COMPUTER WON")
		case user == "scissor" && computer == "rock":
			fmt.Println("COMPUTER WON")
		default:
			fmt.Println("INvalid Chocie")

		}
		

	}

}
