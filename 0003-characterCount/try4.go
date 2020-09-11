package main

import (
	"fmt"
	"strings"
)

func stringCharacterCount(vString string) string {
	vString = strings.ToLower(vString)

	for _, r := range vString {
		fmt.Printf("%c:\n", r)
	}
	return vString
}

// main function
func main() {
	var vString string
	fmt.Println("Enter string to check string character count:")
	fmt.Scanf("%s\n", &vString)
	fmt.Println(stringCharacterCount(vString))
}
