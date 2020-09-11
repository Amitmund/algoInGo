package main

import (
	"fmt"
	"strings"
)

// https://www.geeksforgeeks.org/rune-in-golang/
// The rune type is an alias of int32. Its for Unicode.

func stringCharacterCount(vString string) string {
	fmt.Println("output data: ", vString)
	vString = strings.ToLower(vString)

	// About string length.
	// vlen := len(vString)
	// fmt.Printf("length of the string %s\n", vlen)
	// fmt.Println(vlen)

	// for vChar := 0; vChar <= vlen; vChar++ {
	// 	fmt.Println(vChar)
	// }

	for _, r := range vString {
		//fmt.Printf("i%d r %c\n", i, r)
		fmt.Printf("%c:\n", r)

	}

	//vMap := []rune(vString)

	return vString
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var vString string
	fmt.Println("Enter string to check string character count:")
	fmt.Scanf("%s\n", &vString)
	// To know type of the variable %T
	// fmt.Printf("Type of vString: %T\n", vString)

	fmt.Println(vString, stringCharacterCount(vString))

	// b := []byte(vString)
	// fmt.Printf("Type of b: %T\n", b)
	// fmt.Println(b)

	type point struct {
		x, y int
	}

	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)

	var strs = []string{"T", "E", "S", "T"}

	fmt.Println(Map(strs, strings.ToUpper))
	fmt.Println(Map(strs, strings.ToLower))
}

/*

Links to ref:

https://www.dotnetperls.com/convert-slice-string-go

https://stackoverflow.com/questions/18556693/slice-string-into-letters

*/
