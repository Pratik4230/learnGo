package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while get request", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error ", response.StatusCode)
		return
	}

	// data, errr := ioutil.ReadAll(response.Body)
	// if errr != nil {
	// 	fmt.Println("Error While reading response body", err)
	// 	return
	// }

	// fmt.Println("response body", string(data))

	// Other way

	var todo Todo
	err3 := json.NewDecoder(response.Body).Decode(&todo)
	if err3 != nil {
		fmt.Println("Not able to decode", err3)
		return

	}

	fmt.Println("Todo:", todo)

}
