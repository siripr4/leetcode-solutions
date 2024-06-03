package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

var invertedParenthesisPairs = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := new(Stack)
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == "(" || char == "{" || char == "[" {
			stack.Push(s[i])
		} else if element, isPresent := stack.Pop(); !isPresent || invertedParenthesisPairs[char] != string(element) {
			fmt.Println(string(element))
			fmt.Println(isPresent)
			return false
		}
		fmt.Println(stack)
	}
	return stack.IsEmpty()
}

type Stack []byte

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(element byte) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	lastIndex := len(*s) - 1
	element := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return element, true
}
