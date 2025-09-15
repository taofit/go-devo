package main

import (
	"fmt"
	"strconv"
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
	var str = FormatWithCommas(2928734)
	fmt.Println(numbers)
	fmt.Println(str)

	input := "Gemini"
	fmt.Printf("Hello, %s!\n", input)
}

// FormatWithCommas formats an integer with commas every 3 digits.
func FormatWithCommas(n int) string {
	s := strconv.Itoa(n)
	if len(s) <= 3 {
		return s
	}
	result := ""
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			result = "," + result
		}
		result = string(s[i]) + result
		count++
	}
	return result
}