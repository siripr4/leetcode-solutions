package main

import (
	"strings"
	"unicode"
)

func main() {

}

func isPalindrome(s string) bool {
	f := func(char rune) rune {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return -1
		}
		return unicode.ToLower(char)
	}
	s = strings.Map(f, s)

	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
