package main

import (
	"fmt"
	"os"
	"strconv"

	"addingNNumber/addingN"
)

func main() {

	// Checking argument.

	// TODO: Check the number should be unsinged one.
	// TODO: Create a test check.
	if len(os.Args) != 2 {
		fmt.Printf("How to run \n %s unsignedInt\n\n", os.Args[0])
		fmt.Printf("Example: \n%s 5\n", os.Args[0])
		return
	}

	// Calling function
	number, _ := strconv.Atoi(os.Args[1])
	result := addingN.AddingN(number)
	fmt.Printf("Adding Numbers from 1 to %d is: %d\n", number, result)

}
