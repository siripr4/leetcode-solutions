package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	sToT := make(map[uint16]uint16)
	tToS := make(map[string]string)

	for i := 0; i < len(s); i++ {
		if value, exists := sToT[string(s[i])]; !exists {
			if _, present := tToS[string(t[i])]; present {
				return false
			}
			sToT[string(s[i])] = string(t[i])
			tToS[string(t[i])] = string(s[i])
		} else if value != string(t[i]) {
			return false
		}
	}
	return true
}
