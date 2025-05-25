package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, er := http.NewRequest(http.MethodDelete, myUrl, nil)
	if er != nil {
		fmt.Println("Error new Request", er)
		return
	}

	//send request

	client := http.Client{}
	ress, eru := client.Do(req)

	if eru != nil {
		fmt.Println("Errro sending req", eru)
		return
	}

	defer ress.Body.Close()

	data, _ := ioutil.ReadAll(ress.Body)
	fmt.Println("data", string(data))
	fmt.Println("status", ress.StatusCode)

}
