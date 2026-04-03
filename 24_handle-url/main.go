package main

import (
	"fmt"
	"net/url"
)

func main() {
	// handle url

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	parsedUrl ,err:=url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error at getting url ", err)
		return
	}
	fmt.Println("parsedUrl is :", parsedUrl)

	// get schemems , path , host , rawQuery from the url 
	fmt.Println("host is :", parsedUrl.Host)
	fmt.Println("Path is :", parsedUrl.Path)
	fmt.Println("Scheme is :", parsedUrl.Scheme)
	fmt.Println("RawQuery is :", parsedUrl.RawQuery)

	// update the parsed url custome type
	parsedUrl.Scheme = "ravi"
	parsedUrl.Host = "google.golang.com"
	parsedUrl.Path = "godoc/1"
	parsedUrl.RawQuery = "handle/url/documentation"

	modifiedUrl := parsedUrl.String()

	fmt.Println("new url is :" , modifiedUrl)

}