package main

import (
	"fmt"
	"strings"
)

func main() {
	// concantenate string
	// 1. using + => not very efficient
	h := "Hello, "
	w := "World."

	myString := h + w
	fmt.Println(myString) // Hello, World.
	fmt.Println()

	// 2. using Sprintf => more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString) // Hello, World.
	fmt.Println()

	// 3. using stringbuilder => very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String()) // Hello, World.

	// Getting Substring
	abc := "ABCDEF"
	fmt.Println(abc[0:])  // ABCDEF
	fmt.Println(abc[3:5]) // DE
	fmt.Println(abc[3])   // 68

	// strings
	courses := []string{
		"Learn Go for Beginners",
		"Learn Python for Beginners",
		"Learn JavaScript for Beginners",
		"Learn C# for Beginners",
	}

	// Contains
	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in word:", x, "and index is", strings.Index(x, "Go"))
		}
	}

	// HasPrefix
	newString := "Go is a great programming language, and Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go")) // true
	fmt.Println(strings.HasSuffix(newString, "Go")) // false
	fmt.Println(strings.Count(newString, "Go"))     // 2
	fmt.Println(strings.LastIndex(newString, "Go")) // 40
	fmt.Println(strings.Index(newString, "Go"))     // 0

	// if strings.Contains(newString, "Go") {
	newString = strings.Replace(newString, "Go", "Golang", -1)
	// }

	fmt.Println(newString)

}
