package main

import (
	"fmt"
)

const user = "admin"

func main() {
	numbers := map[string]int{
		"ki":  834,
		"kis": 834,
		"ksi": 834,
		"ski": 834,
	}
	numbers["ssk"] = 8348
	for index, value := range numbers {
		fmt.Println(index, value)
	}
	fmt.Println(numbers)
}
