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

	buffer := 0
	previous := 0
	sum := 0

	for index, val := range s {
		char := string(val)
		curr := symbolValues[char]
		if index == 0 || curr <= previous {
			fmt.Print("if ")
			sum = sum + curr - buffer
			buffer = 0
			fmt.Printf("previous: %d, curr: %d, buffer: %d, sum: %d\n", previous, curr, buffer, sum)
		} else {
			fmt.Print("else ")
			buffer = buffer + previous
			// sum = sum + curr
			fmt.Printf("previous: %d, curr: %d, buffer: %d, sum: %d\n", previous, curr, buffer, sum)
		}
		previous = curr
	}
	return sum
}
