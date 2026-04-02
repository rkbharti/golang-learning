// split,count, joinn, remove(trim) whitespace
// csv type analyser

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the Inputs with ',' : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading INput")
		return 
	}
	// fmt.Println(name)

	parts := strings.Split(input, ",")

	fmt.Println("here the name as we given by input")
	for i, value := range parts {
		parts[i] = strings.TrimSpace(value)

		
		fmt.Println(parts[i])
		
	}
	// count
	count := make(map[string]int)
	for _, name := range parts{
		name = strings.ToLower(name)
		if name == ""{
			continue
		}
		count[name]++
	}

	// display the output
	for name , freq := range count{
		fmt.Println("here is the total no value for each element")
		fmt.Println(name , "--", freq)
	}

}
