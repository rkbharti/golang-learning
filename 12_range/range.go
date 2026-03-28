// used for iteration over data strutures and mostly used

package main

import "fmt"

func main(){
	// nums :=[]int{5,6,7}
	// sum := 0
	// for k, num:=range(nums){
	// 	fmt.Println(num, k)
	// 	sum = sum + num
	// }
	// fmt.Println("sum is : ", sum)

	// map in range for string
	m:=map[string]string{"fname" : "ravi" , "lname": "kumar"}
	  for k , v := range m{
		fmt.Println(k)
		fmt.Println( v)
	  }

	for k , v := range "go"{
		
		fmt.Println("this will give key and unicde for the given string", k, v)
		fmt.Println(k, string(v))
		
	}
	

}

