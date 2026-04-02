package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// we implemt os.Create to create a nmed text file and use defer to close it again after all function run
	// file, err := os.Create("xyz.txt")
	// if err != nil {
	// 	fmt.Println("error at creating file ", err)
	// 	return
	// }
	// defer file.Close()

	// fmt.Println("file created succesfully")

	// // write something into the file
	// content := "HEllo ! whtsapp GUyzzz"

	// _, err2 := io.WriteString(file, content)
	// if err2 != nil {
	// 	fmt.Println("error at writing", err2)
	// 	return
	// }
	// fmt.Println("content added succesfully")

	// open and read from the file using bufio
	file3, err := os.Open("xyz.txt")
	if err != nil {
		fmt.Println("error at opeing file ", err)
		return
	}
	defer file3.Close()

	scanner := bufio.NewScanner(file3)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
