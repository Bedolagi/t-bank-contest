package main

import (
	"fmt"
	"slices"
	"strings"
)

func isAlmostPalindrome(example string) string {
	for i := 0; i < 4; i++ {
		candidate := example[:i] + example[i+1:]
		chars := strings.Split(candidate, "")
		slices.Reverse(chars)
		reversed := strings.Join(chars, "")
		if candidate == reversed {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	var example string
	fmt.Scanln(&example)
	fmt.Println(isAlmostPalindrome(example))
}
