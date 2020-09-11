package main

import (
	"fmt"
	"sort"
)

func main() {
	strSlice := []string{"Jamaica", "Estonia", "Indonesia", "Hong Kong"} // unsorted
	fmt.Println("Slice of string BEFORE sort:", strSlice)
	sort.Strings(strSlice)
	fmt.Println("Slice of string AFTER  sort:", strSlice)

	fmt.Println("\n-----------------------------------\n")

	strSlice = []string{"JAMAICA", "Estonia", "indonesia", "hong Kong"} // unsorted
	fmt.Println("Slice of string BEFORE sort:", strSlice)
	sort.Strings(strSlice)
	fmt.Println("Slice of string AFTER  sort:", strSlice)
}
