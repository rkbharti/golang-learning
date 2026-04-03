package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Details struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Country  string `json:"country"`
}

func main() {
	// parseurl using http.Get

	res, err := http.Get("https://fake-json-api.mock.beeceptor.com/users/1")
	if err != nil {
		fmt.Println("error at recieing url response", err)
		return
	}
	defer res.Body.Close()

	// we want this data in json.Marshal
	var data Details
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("error at marshal json", err)
		return
	}
	data.Country = "India"
	fmt.Printf("ID: %d\nName: %s\nUsername: %s\nCountry: %s\n",
		data.Id, data.Name, data.Username, data.Country)

	// json marshal
	jsondata, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error is at encoding", err)
		return
	}
	fmt.Println("json data of !st id is :", string(jsondata))

}
