package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://fakestoreapi.com/products")

	if err != nil {
		fmt.Println("Error white get ", err)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Type of Response is %T\n", res)

	//Read ResponsBody

	var data, errr = ioutil.ReadAll(res.Body)
	if errr != nil {
		fmt.Println("Error While reading Response body", errr)
		return
	}

	fmt.Println("Response Body is", string(data))

}
