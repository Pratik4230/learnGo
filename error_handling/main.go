package main

import "fmt"

func multy(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator should not zero")
	}

	return a / b, nil
}

func main() {
	ans, _ := multy(5, 0)
	fmt.Println("Answer is", ans)
}
