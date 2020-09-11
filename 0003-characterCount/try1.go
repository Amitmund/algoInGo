package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a long string with many repeated characters"
	numberOfa := strings.Count(str, "a")

	fmt.Printf("[%v] string has %d of characters of [a] ", str, numberOfa)

}
