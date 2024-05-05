package main

import "fmt"

// Problem: [Roman to Integer](https://leetcode.com/problems/roman-to-integer/description)
func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	symbolValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	previous := 0
	sum := 0

	for i := len(s) - 1; i >= 0; i-- {
		curr := symbolValues[string(s[i])]
		if curr < previous {
			sum -= curr
		} else {
			sum += curr
		}
		previous = curr
	}
	return sum
}
