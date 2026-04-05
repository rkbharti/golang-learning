package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Employee struct {
	ID       int    `json:"id"`
	Fname    string `json:"firstName"`
	Lname    string `json:"lastName"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	const myUrl = "https://dummyjson.com/users/58"
	req, err := http.NewRequest(http.MethodGet, myUrl, nil)
	if err != nil {
		fmt.Println("error at getting reponse", err)
		return
	}

	// create client request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error at client do", err)
		return
	}
	defer res.Body.Close()

	// read the raw bytes  fromm the  reponse body using readAll
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err at reading raw bytes", err)
		return
	}

	// create an empty instance for the object
	var user Employee

	// unmarshal the bytes we got using ReadAll
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		fmt.Println("error at unmarshal ", err)
		return
	}

	fmt.Println("User Object CReated succcesfully")
	fmt.Println("Id is : ", user.ID)
	fmt.Println("Fnam is :", user.Fname)
	fmt.Println("Lname is :", user.Lname)
	fmt.Println("Username  is :", user.Username)
	fmt.Println("Password  is :", user.Password)
}
