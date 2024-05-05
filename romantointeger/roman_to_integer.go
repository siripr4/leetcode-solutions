package main

import "fmt"

// Problem: [Roman to Integer](https://leetcode.com/problems/roman-to-integer/description)
func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	symbolValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	buffer := 0
	previous := 0
	sum := 0

	for _, char := range s {
		curr := symbolValues[char]
		if curr <= previous {
			sum = sum + previous - buffer
			buffer = 0
		} else {
			buffer = buffer + previous
		}
		previous = curr
	}
	sum = sum - buffer + previous
	return sum
}
