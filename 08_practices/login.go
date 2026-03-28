// CLI Login System + Menu

package main

import (
	"fmt"
)
const(
		ADMIN_ROLE = "admin"
		ADMIN_PASSWORD = "1234"

		USER_ROLE = "user"
		USER_PASSWORD = "7890"
	)

func main() {
	var role string
	var  password string
	// LOGIN
	fmt.Println("enter ur role = " )
	fmt.Scanln(&role)
	fmt.Println("enter ur password = " )
	fmt.Scanln(&password)

	// AUTH CHECK
	if (role == ADMIN_ROLE && password == ADMIN_PASSWORD)  ||
		(role ==USER_ROLE && password == USER_PASSWORD){
			fmt.Println("LOGIN SUCCESSFULLY")

			if role == ADMIN_ROLE{
				fmt.Println("Full Acccess Granted")
			} else {
				fmt.Println("Limited acces granted")
			}

			// menu loop
		for {
			var choice int
			fmt.Println("\n -----MENU-----")
			fmt.Println("1. WELCOME ")
			fmt.Println("2. PRINT NUMBER 1 TO 10 ")
			fmt.Println("3. Multiplication of ur desired number ")
			fmt.Println("4. Check Even/Odd Number ")
			fmt.Println("5. EXIT")
			fmt.Println("Enter you choice")
			fmt.Scanln(&choice)

			switch choice{
			case 1:
				fmt.Println("WELCOME TO DASHBOARD",role)
			case 2:
				for i := 0;i<=10;i++{
					fmt.Println(i)
				}
			case 3:
				var num int
				fmt.Println("enter ur no for multiplication")
				fmt.Scanln(&num)
				for i :=1;i<=10;i++{
					fmt.Println(num, "x", i, "=", num*i)
				}
			case 4:
				var num2 int
				fmt.Println("Enter the NO to Check ODD/EVEN: ")
				fmt.Scanln(&num2)
				if num2 % 2 == 0{
					fmt.Println("EVEN NO")
				}else{
					fmt.Println("ODD NO")
				}

			
			case 5:
				fmt.Println("EXISTING ")
				return
			default:
				fmt.Println("WRONG INPUT")
			}

		}
	} else {
		fmt.Println("INVALID CREDENTIALS")
	}
}