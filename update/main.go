package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	todo := Todo{
		UserId:    32,
		Title:     "title",
		Completed: true,
	}

	//convert to json

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Errror coverting to json", err)
		return
	}

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, er := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if er != nil {
		fmt.Println("Error new Request", er)
		return
	}

	req.Header.Set("Content-type", "application/json")

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

}
