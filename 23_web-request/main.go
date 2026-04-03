package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// get a http request from an url
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error at getting request ", err)
		return
	}
	defer res.Body.Close()

	// fmt.Println("response is :" , res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error at reading data ", err)
		return
	}
	fmt.Println("reposne of given url is :", string(data))
}
