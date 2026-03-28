// function are first class citizen which means we can assign function to any variable
// or , we can pas one function to another as argument or we can return a new function from another function
package main

import "fmt"

func add(a, b int) int {
	return a + b // function to add two number
}

func languages()(string, string,string , bool){
	return "hii","from" , "golang", false //function to print multiple string 
}
 
func processit (fn func(a int)int){
	fn(2)
}

func processit2 ()func(a int)int{
	return  func(a int) int {
		return 4
	}
}
func main() {
	total := add(5, 6)
	fmt.Println(total) //sum
	greet1 , greet2, greet3, _ :=languages()
	fmt.Println(greet1, greet2, greet3) //return string or  bool

	fn := func(a int)int  {
		return 2
	}
	processit(fn)

	fn2 := processit2()
	fn2(9)

		
}