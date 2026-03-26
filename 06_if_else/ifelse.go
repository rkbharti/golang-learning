package main

import "fmt"

func main() {
	// using if-else-if to print age as ADULT, TEENAGER, KID
	// age := 10

	// if age >= 18 {
	// 	fmt.Println(age, "age belong to  Adult")
	// } else if age >= 11 {
	// 	fmt.Println(age ,"age belong to  Teenager")
	// } else {
	// 	fmt.Println(age,"age belong to kid ")
	// }

	// NOW in this we  can use operator  and constants grouping  in if else for condition checking
	const(
		ADMIN_ROLE = "admin"
		ADMIN_PASSWORD = "1234"

		USER_ROLE = "user"
		USER_PASSWORD = "7890"
	)
	var role string
	var password string

	fmt.Println("enter role: ")
	fmt.Scanln(&role)
	fmt.Println("------------------------------------------------------")
	fmt.Println("enter your password")
	fmt.Scanln(&password)

	if role == ADMIN_ROLE && password == ADMIN_PASSWORD {
		fmt.Println("admin accessed")
	} else if role == USER_ROLE && password == USER_PASSWORD{
		fmt.Println("User accessed")
	} else {
		fmt.Println("wrong credentials")
	}
}
