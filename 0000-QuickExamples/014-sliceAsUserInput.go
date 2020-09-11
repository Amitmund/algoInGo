// Go program to illustrate how to pass a slice to the function
package main

import "fmt"

// Function in which slice is passed by value
func myfun(element []string) {

	// Modifying the given slice
	element[2] = "Java"
	fmt.Println("Modified slice: ", element)
}

// main function
func main() {
	var vString string
	fmt.Println("Enter a slice, Example: {\"C#\", \"Python\", \"C\", \"Perl\"} ")
	// fmt.Scanf("%s\n", &slc)
	fmt.Scanf([]string &slc)
	fmt.Println("Initial slice: ", slc)

	// Passing the slice to the function
	myfun(slc)

	// fmt.Println(myfun(slc))
	fmt.Println("Final slice:", slc)

}
