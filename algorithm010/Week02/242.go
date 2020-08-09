package main

import (
	"fmt"
)

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var buket [26]int
	for i := range s {
		buket[s[i]-'a']++
		buket[t[i]-'a']--
	}
	for i := range buket {
		if buket[i] != 0 {
			return false
		}
	}
	return true
}
