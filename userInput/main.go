package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hey Enter your full name")

	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Your name is", name)

	reader := bufio.NewReader(os.Stdin)
	newName, _ := reader.ReadString('\n')
	fmt.Println("Hello ", newName)

}
