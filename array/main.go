package main

import "fmt"

func main() {

	var name [5]int

	name[0] = 10
	name[1] = 15
	name[3] = 20

	fmt.Println(name)

	var numbers = [5]int{0, 12, 3, 4, 5}

	fmt.Println(numbers)

	fmt.Println("Length is ", len(numbers))

}
