// variadic function is basically fmt.println(whatvver we type is as many as we want)

package main

import "fmt"

// heree we create a function which help in sum and we can eneter as many values as want
func add(nums ...int)int{
	total:= 0
	for _,num:=range nums{
		total = total + num
	}
	return total

}
func main(){
	m:=[]int{12,2,5,7}
	result := add(m...)
	fmt.Println(result)


}