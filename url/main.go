package main

import (
	"fmt"
	"net/url"
)

func main() {

	myUrl := "https://fakestoreapi.com/products?value=1&key=2"
	fmt.Printf("Type of myUrl is %T \n", myUrl)

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error while Parsing myUrl", err)
		return
	}

	fmt.Printf("Type of URL %T\n", parsedUrl)

	fmt.Println("Scheme : ", parsedUrl.Scheme)
	fmt.Println("Host : ", parsedUrl.Host)
	fmt.Println("Path : ", parsedUrl.Path)
	fmt.Println("RawQuery : ", parsedUrl.RawQuery)

	parsedUrl.RawQuery = "username=abc"

	newUrl := parsedUrl.String()

	fmt.Println("New url is", newUrl)

}
