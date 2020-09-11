package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := []rune(s)
	mid := len(r) / 2
	last := len(r) - 1
	for i := 0; i < mid; i++ {
		if r[i] != r[last-i] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &s)
	fmt.Println(s, isPalindrome(s))
}
