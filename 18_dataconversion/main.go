// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// change from num to float
// 	var num int = 20
// 	fmt.Println("number is " ,  num)
// 	fmt.Printf("data type of num is %T\n", num)

// 	var numtofloat float64 = float64(num)
// 	numtofloat = numtofloat + 10.2
// 	fmt.Println("num change to float value", numtofloat)
// 	fmt.Printf("data type of num is %T\n", numtofloat)

// 	// num change to string
// 	var num1 int = 256
// 	fmt.Println("number is " ,  num1)
// 	fmt.Printf("data type of num is %T\n", num1)
// 	var numtostring string = strconv.Itoa(num1)
// 	numtostring = numtostring + "abcd"
// 	fmt.Println("num1 change to string value", numtostring)
// 	fmt.Printf("data type of num1 is %T\n", numtostring)

// 	// change string to num
// 	var name string = "Ravi"
// 	stringtonum,err := strconv.Atoi(name)
// 	if err != nil{
// 		fmt.Println("NO conversion",err)
// 	}else{
// 		fmt.Println("conversion is ", stringtonum)
// 	}
// 	stringtonum = stringtonum + 2
// 	fmt.Println("name change to int data type", stringtonum)
// 	fmt.Printf("data type of stringtonum is %T\n", stringtonum)

// 	// change string to bool
// 	var name2 string = "false"
// 	fmt.Println("name2  is string now :", name2)
// 	fmt.Printf("data type of num2 is %T\n", name2)
// 	name2tobool,_ := strconv.ParseBool(name2)
// 	fmt.Println("name2change to bool :", name2tobool)
// 	fmt.Printf("data type of num2tobool is %T\n", name2tobool)

// }

package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var a string
		var b string
		fmt.Println("enter the first value")
		fmt.Scanln(&a)
		fmt.Println("Enter the second value")
		fmt.Scanln(&b)

		// string to int /float
		c, err1 := strconv.ParseFloat(a, 64)
		d, err2 := strconv.ParseFloat(b, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input")
			continue
		}

		var choice int
		fmt.Println("----------MENU---------- ")
		fmt.Println("1 : Add")
		fmt.Println("2: subtract")
		fmt.Println("3 : multiply")
		fmt.Println("4 : divide")
		fmt.Println("5:  Exit  ")
		fmt.Println("Enter your choice: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Add is : ", c+d)
		case 2:
			fmt.Println("Subtract  is : ", c-d)
		case 3:
			fmt.Println("Multiplication is ", c*d)
		case 4:
			if d == 0 {
				fmt.Println("Not divide by zero")
			} else {
				fmt.Println("division is :", c/d)
			}

		case 5:
			fmt.Println("existing...")
			return
		default:
			fmt.Println("INvalid choice")
		}
	}

}
