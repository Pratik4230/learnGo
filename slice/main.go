package main

import "fmt"

func main() {
	numbers := []int{1, 23, 4, 4, 5}
	fmt.Println("Numbers ", numbers)

	numbers = append(numbers, 9, 9, 9, 9, 9999)

	fmt.Println("After Append", numbers)

	var days = make([]string, 2, 3)
	// type length capacity
	fmt.Println("days:", days, "length:", len(days), "capacity:", cap(days))

	days = append(days, "qwer", "ddd")
	fmt.Println("days:", days, "length:", len(days), "capacity:", cap(days))

	days = append(days, "iiu", "jniu", "bhh", "hi")
	fmt.Println("days:", days, "length:", len(days), "capacity:", cap(days))

}
