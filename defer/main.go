package main

import "fmt"

func Wildcard() {
	fmt.Println("Wildcard")
}
func main() {

	defer Wildcard()
	fmt.Println("First")
	defer fmt.Println("second")
	fmt.Println("Last")

}
