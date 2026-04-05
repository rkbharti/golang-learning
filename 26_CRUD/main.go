package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

// create function to get data from url and unmarshal it to struct and print it
func getDataFromURL() {
	fmt.Println("use GET method to get data from Url and then marshal the data in obj form to get data from json format to str format for each user ")
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

// create function to post data to url using http.Post and marshal it to json
func postDataToURL() {
	person1 := Details{
		Id:       1,
		Name:     "Ravi",
		Username: "rkbharti",
		Country:  "India",
	}
	// as we created person1 data field with the help of struct we created above now e have to pass this data to json
	// so we have to marshal this data so that json can acces this data
	fmt.Println("we immplement POST method to send data to a specicif url using json.marshal , reader and string because json need string format data to accept ")
	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("error at marshal data ", err)
		return
	}
	// now we to the marshal part of person 1 we have to convert this into string because json accept string format for getiing and postong data
	jsonStr := string(jsonData)

	// now we have to use http/post mehtod to send request to the url and it accept three parameter first is url , second is content , and thirs is reader
	myUrl := "https://jsonplaceholder.typicode.com/todos"
	jsonReader := strings.NewReader(jsonStr)
	jsonResponse, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error at posting req", err)
		return
	}
	defer jsonResponse.Body.Close()
	fmt.Println("Response is : ", jsonResponse.Status)

}

func updateDataToUrl() {
	person1 := Details{
		Id:      1,
		Name:    "Ravi Kumar",
		Country: "UK/US",
	}
	// do the marshal of person 1
	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error at Marshal ", err)
		return
	}
	//new request accept three thing as input method, url , reader
	reader := strings.NewReader(string(jsonData))
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	// now we create PUT request
	req, err := http.NewRequest(http.MethodPut, myUrl, reader)
	if err != nil {
		fmt.Println("error at request", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// now we have to send the request using client
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error at send request by client", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error at Reading", err)
		return
	}
	fmt.Println("Response Status is :", res.Status)
	fmt.Println("Update Response is :", string(data))

}

func deleteDataToUrl() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"
	// create delete request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("error at New REQUEST ", err)
		return
	}
	// create client side request

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error at client Do ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("STATUS OF DELETION IS :", res.Status)

}

func main() {
	getDataFromURL()
	postDataToURL()
	updateDataToUrl()
	deleteDataToUrl()

}
