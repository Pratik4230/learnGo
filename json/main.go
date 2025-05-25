package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsPassed bool   `json:"isPassed"`
}

func main() {

	person := Person{Name: "Pratik", Age: 22, IsPassed: true}
	fmt.Println("person data is ", person)

	//convert person into json / encoding / Marshling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshlling", err)
		return
	}

	fmt.Println("Person Data is", string(jsonData))

	//decoding Unmarshlling

	var personData Person

	errr := json.Unmarshal(jsonData, &personData)
	if errr != nil {
		fmt.Println("Error Unmarshling", errr)
		return
	}

	fmt.Println("Person Data after decoding", personData)

}
