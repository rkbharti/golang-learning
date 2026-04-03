package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// create a file
	// file, err := os.Create("abc.txt")
	// if err != nil {
	// 	fmt.Println("err at creating file ")
	// 	return
	// }

	// // write in file 
	// content := "HEy this is RAVI "
	// _,errors:=io.WriteString(file, content)
	// if errors != nil {
	// 	fmt.Println("error at writing in file ", errors)
	// 	return
	// }

	// fmt.Println("file created succesfullly")

	// defer file.Close()

	// read from file 

	file, err:=os.Open("abc.txt")
	if err != nil {
		fmt.Println("err at creating file ", err)
		return
	}
	defer file.Close()

	// read file using buffer
	buffer := make([]byte , 1024)

	for {
		n, err:=file.Read(buffer)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("error at readding file ", err)
			return 
		}
		fmt.Println("file contains in abc.txt is: ", string(buffer[:n]))


	}
	


}
