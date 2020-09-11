// https://stackoverflow.com/questions/41690156/how-to-get-the-keys-as-string-array-from-map
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	a := map[string]int{
		"A": 1, "B": 2,
	}
	keys := reflect.ValueOf(a).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	fmt.Print(strings.Join(strkeys, ","))
}
