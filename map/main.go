package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Pratik"] = 100
	studentGrades["Pratik2"] = 99
	studentGrades["Pratik3"] = 98

	fmt.Println("Marks of Pratik is", studentGrades["Pratik"])

	studentGrades["Pratik"] = 101

	fmt.Println("Updated ", studentGrades["Pratik"])

	delete(studentGrades, "Pratik3")

	marks, isExists := studentGrades["Pratik3"]

	fmt.Println("Grades of Pratik3 ", marks)
	fmt.Println("but Pratik3  exists ? ", isExists)

	for index, value := range studentGrades {
		fmt.Printf("The Index is %s  & value is %d\n", index, value)
	}

	///direct
	person := map[string]int{
		"Pratik":  20,
		"Pratik2": 21,
		"Pratik3": 22,
	}

	for index, value := range person {
		fmt.Printf("The Index is %s  & value is %d\n", index, value)

	}
}
