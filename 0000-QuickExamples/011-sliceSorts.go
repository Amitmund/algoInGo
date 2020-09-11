package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{10, 5, 25, 351, 14, 9} // unsorted
	fmt.Println("Slice of integer BEFORE sort:", intSlice)
	sort.Ints(intSlice)
	fmt.Println("Slice of integer AFTER  sort:", intSlice)
}
