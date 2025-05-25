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

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, errr := http.Post(myUrl, "application/json", jsonReader)
	if errr != nil {
		fmt.Println("Error Post request", errr)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("data", string(data))

}
