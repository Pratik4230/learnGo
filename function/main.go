package main

import "fmt"

func simple() {
	fmt.Println("Welcome to Functions")
}

func add(a, b int) int {
	return a + b
}

func multy(a int, b int) (result int) {
	result = a * b

	return
}

func main() {
	simple()
	fmt.Println(add(5, 5))
	fmt.Println(multy(5, 5))
}
