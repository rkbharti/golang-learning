// slices are basically dynamic array where we dont know how much size or data is going to store and we use this construct mosstly

package main

import (
	"fmt"
	"slices"
)

func main(){
	// by default uninitailsed slice is nil 
	fmt.Println("default uninitailsed slice is nil ")
	var num[]int
	fmt.Println(num == nil) //output:true
	fmt.Println(len(num)) //in-built function to fin the length of array

	// to make it initialize and remove it default behvaiour we use make() function 
	fmt.Println("here we make this initialize using maake() function")
	var num2 = make([]int,1,7) // here in make([this is empty array]data-type, 3-initialise size, 7- capacity (max no of element and it change when more elemts addded))
	num2 =append(num2, 1 , 2 , 3 , 4 , 5 , 6 , 7) // append is in-built function to add element in slice and it return new slice with added element
	// append no at index 
	num2[0]= 3
	num2[1] = 5 //it can overwrite data
	fmt.Println(num2)
	fmt.Println(cap(num2))
	fmt.Println(len(num2)) // in-built function to find the length of slice

	// add elemt in single line while declaring the slice
	fmt.Println("here we learn how to add element in single line")
	num3:=[]int{}
	num3 = append(num3, 1,2,23,56,25, 20, 10 )
	fmt.Println(num3)
	fmt.Println("capacity of array increase when we put more element into the list",cap(num3))
	fmt.Println("length of array",len(num3))


	fmt.Println("here we learn how to  copy element ")
	copy1:=[]int{}
	copy1 = append(copy1, 5)

	var copy2 = make([]int,len(copy1))

	
	// copy function
	copy(copy2, copy1)
	fmt.Println(copy1 , copy2)


	// Slice operator
	fmt.Println("using slice operator")

	var age = []int{20,21,22}
	fmt.Println(age[0:1]) //exlude last index [20]
	fmt.Println(age[:2]) // [20,21]
	fmt.Println(age[0:]) // [20,21,22]

	// slice package
	fmt.Println("using Slice package")
	var quantity = []int{3,4}
	var quanitity2 = []int{3,4}
	fmt.Println(slices.Equal(quantity,quanitity2)) //true

	//making two D slices
	fmt.Println("using 2D slices")
	var twoDArray = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(twoDArray)
}