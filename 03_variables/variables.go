package main

import "fmt"

func main() {
	// declare string variable
	var name string = "Hello from Variable"
	var name2 = "Hello  by infer mean go can identify which variable we are using "
	// using shorthand syntax
	name3 := "hello by shorthand syntax using :="
	// normal syntax
	var age int=25
	// using infer where it automatically find out 
	var age2 = 20
	// using shorthand syntax
	age3 := 29
	// for bool values we can use three types to print this according to need like when and where we need to assign the value later or earlier 
	var normalSyntax bool = true 
	var isTrue = true
	isFalse := false
	// printing all values
	fmt.Println(name) 			//using normal syntax
	fmt.Println(name2)  		//using infer where go identify it by self
	fmt.Println(name3)  		// using := for maiing it more shorthand syntax
	fmt.Println(age)			//using normal syntax
	fmt.Println(age2)			//using infer where go identify it by self
	fmt.Println(age3)			// using := for maiing it more shorthand syntax
	fmt.Println(normalSyntax) 	//using normal syntax
	fmt.Println(isTrue)  		//using infer where go identify it by self
	fmt.Println(isFalse) 		// using := for maiing it more shorthand syntax
	
	
}