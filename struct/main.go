package main

import "fmt"

type Person struct {
	Name  string
	age   int
	marks int
}

type Contact struct {
	email string
	phone int
}

type Employee struct {
	EmpBasic   Person
	EmpContact Contact
}

func main() {
	var Pratik Person

	fmt.Println("Pratik", Pratik)

	Pratik.Name = "Praik"
	Pratik.age = 20
	Pratik.marks = 100

	fmt.Println("Pratik : ", Pratik)

	// 2nd

	Pratik2 := Person{
		Name:  "Pratik2",
		age:   21,
		marks: 90,
	}

	fmt.Println("Pratik 2 is : ", Pratik2)

	//3rd
	var Pratik3 = new(Person)
	Pratik3.Name = "Pratik3"
	Pratik3.age = 22
	Pratik3.marks = 80

	fmt.Println("Pratik3 ", Pratik3)

	var pratik4 Employee
	pratik4.EmpBasic = Person{
		Name:  "Pratik4",
		age:   25,
		marks: 85,
	}

	pratik4.EmpContact = Contact{
		email: "Pratik4@gmail.com",
		phone: 987654345,
	}

	fmt.Println("Person4", pratik4)

}
