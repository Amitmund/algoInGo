package main

import (
	"fmt"
)

// main function
func main() {
	var vString string
	fmt.Println("Enter string to check string character count:")
	fmt.Scanf("%s\n", &vString)

	for i := 0; i < len(vString); i++ {
		fmt.Println(string(vString[i]))
	}

}
