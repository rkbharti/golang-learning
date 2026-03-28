package main

import "fmt"

func main() {
	var nums[3]int
	fmt.Println(len(nums)) //to get length of array
	//add element in array
	nums[0]=1
	fmt.Println(nums[0]) // print the index number
	fmt.Println(nums) //print all number store in array

	var num[2]bool
	num[1]= true // zeroth elemt will depend upon type for bool with nil= false , string =[], int=0
	fmt.Println(num)

	var name[3]string
	name[2]="golang"
	fmt.Println(name)

	// add number when we declre the array in single line

	age:=[3]int{
		20,18,25}
		fmt.Println(age)

	//two dimensional array
	twoDArray:=[3][2]int{{2,3},{4,5}}
	fmt.Println(twoDArray)
	// when we knwo we have fixed size , data and help in memory optimixation

}