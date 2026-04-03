// json - javascript object notation

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	IsAdult bool   `json:"isadult"`
}

func main() {
	
	person1 := Person{
		Name: "Ravi",
		Age: 25,
		IsAdult: true,
	}

	fmt.Println("person1 data is :" , person1)

	// convert person into json marshal (encoding) using package encoding/json
	jsonData, err :=json.Marshal(person1)
	if err != nil {
		fmt.Println("error at marshaling data ", err)
	}
	fmt.Println("person1 data after marshal is :" , string(jsonData))


	//now if we want to demarshal this person data 
	var personData Person
	errors  := json.Unmarshal(jsonData,&personData)
	if errors != nil {
		fmt.Println("error at unmarshaling", errors)
	}
	fmt.Println("unmarshal data is :", personData)



}
