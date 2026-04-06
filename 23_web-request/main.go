// use crud operation for web request

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myUrl = "https://dummyjson.com/users/55"

// struct Employee
type Employee struct {
	Id       int    `json:"id"`
	Fname    string `json:"firstName"`
	Lname    string `json:"lastName"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func styler() {
	fmt.Println("**********************************************")
}

// create a function for get method from url
func getMethodFromUrl() {
	req, err := http.NewRequest(http.MethodGet, myUrl, nil)
	if err != nil {
		fmt.Println("error at getting data ", err)
		return
	}
	// create client side request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error at getting response", err)
		return
	}
	defer res.Body.Close()

	// Read byte data from response body
	jsondata, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error at Reading data", err)
		return
	}

	// here we create a variable in which we want our data to store
	var user Employee

	// now we do the demarhal from json format
	err = json.Unmarshal(jsondata, &user)
	if err != nil {
		fmt.Println("error at unmarshal", err)
		return
	}
	fmt.Println("staus of gettinng data is", res.Status)
	fmt.Println("Data fetch succesfully")
	fmt.Println("Id of the user is :", user.Id)
	fmt.Println("fullName of user is :", user.Fname, user.Lname)
	fmt.Println("Credential details are :", "Username is :", user.Username, "Password is ", user.Password)

}

// create function for post method to send data to url
func postMethodFromUrl() {
	person1 := Employee{
		Id:       55,
		Fname:    "Ravi",
		Lname:    "Kumar",
		Username: "rkbharti",
		Password: "rkb2002",
	}
	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("error at marshal ", err)
	}
	// now we have to convert jsonData into  string so that json accept this data to post into server
	jsonStr := string(jsonData)

	// create post REQUEST
	const myUrl = "https://dummyjson.com/users/add"
	reader := strings.NewReader(jsonStr)

	req, err := http.NewRequest(http.MethodPost, myUrl, reader)
	if err != nil {
		fmt.Println("error at NEwREquest", err)
		return
	}
	// tell the server u are sending Json
	req.Header.Set("Content-Type", "application/json")

	// create cleint side request
	client := http.Client{}
	data, err := client.Do(req)
	if err != nil {
		fmt.Println("error at client side ", err)
		return
	}
	defer data.Body.Close()
	fmt.Println("respose staus is  :", data.Status)
	res, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Println("error at reading ", err)
		return
	}
	fmt.Println("response for post is :", string(res))
}

func putMethodFromUrl() {
	person1 := Employee{
		Id:       26,
		Fname:    "Ravi",
		Lname:    "Kumar Bharti",
		Username: "rkb",
		Password: "rkb20000",
	}
	// rist we have to marshal this dat to send to the web req from the url
	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("error at marshal ", err)
		return
	}
	// we have to convert the jsondata to string so that we can pass it to json
	jsonStr := string(jsonData)

	// create PUT request
	reader := strings.NewReader(jsonStr)
	req, err := http.NewRequest(http.MethodPut, myUrl, reader)
	if err != nil {
		fmt.Println("error at newRequest", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// create client side request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("errorr at client request", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Update Status is :", res.Status)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error at Reading data", err)
		return
	}
	fmt.Println("Update data is :", string(data))
}

func deleteMethodFromUrl() {
	// create delete request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("error at NewRequest", err)
		return
	}

	// create client side request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error at client request", err)
	}
	defer res.Body.Close()
	fmt.Println("delete status is :", res.Status)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error at reading data ", err)
		return
	}
	fmt.Println("deleted data is :", data)

}

func main() {
	styler()
	fmt.Println("Get Method using url")
	getMethodFromUrl()
	styler()
	fmt.Println("Post method from url ")
	postMethodFromUrl()
	styler()
	fmt.Println("Put request to update the value")
	putMethodFromUrl()
	styler()
	fmt.Println("Delete method from the url ")
	deleteMethodFromUrl()
	styler()

}
