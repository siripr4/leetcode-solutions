package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	var prefix string
	for i := 0; ; i++ {
		var currChar string
		for index, str := range strs {
			if i < len(str) {
				if index == 0 {
					currChar = string(str[i])
				} else if currChar != string(str[i]) {
					return prefix
				}
			} else {
				return prefix
			}
		}
		if currChar == "" {
			break
		}
		prefix = prefix + currChar
	}
	return prefix
}
