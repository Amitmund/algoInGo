// https://golang.org/pkg/strings/

package main

import (
	"fmt"
	"strings"
)

func main() {
	strSlice := strings.SplitN("a,b,c", ",", 0)
	fmt.Println(strSlice, "\n")
	fmt.Printf("Type: %T\n", strSlice)

	strSlice = strings.SplitN("a,b,c", ",", 1)
	fmt.Println(strSlice, "\n")
	fmt.Printf("Type: %T\n", strSlice)

	strSlice = strings.SplitN("a,b,c", ",", 2)
	fmt.Println(strSlice, "\n")
	fmt.Printf("Type: %T\n", strSlice)

	strSlice = strings.SplitN("a,b,c", ",", 3)
	fmt.Println(strSlice, "\n")
	fmt.Printf("Type: %T\n", strSlice)

}
