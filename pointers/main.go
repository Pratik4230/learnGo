package main

import "fmt"

func main() {

	var num int
	num = 2

	var ptr *int
	ptr = &num

	fmt.Println("Pointer containes : ", ptr)
	fmt.Println("Data container at Pointer is : ", *ptr)

}
