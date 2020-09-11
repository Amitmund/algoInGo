// https://www.dotnetperls.com/convert-slice-string-go
package main

import (
	"fmt"
	"strings"
)

func main() {
	values := []string{"one", "two", "three"}

	// Convert string slice to string.
	// ... Has comma in between strings.
	result1 := strings.Join(values, ",")
	fmt.Println(result1)

	// ... Use no separator.
	result2 := strings.Join(values, "")
	fmt.Println(result2)
}

/*
Output:

one,two,three
onetwothree
*/
