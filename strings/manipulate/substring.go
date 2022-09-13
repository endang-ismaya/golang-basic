package manipulate

import (
	"fmt"
	"strings"
)

func GetSubstring() {
	// Getting Substring
	abc := "ABCDEF"
	fmt.Println(abc[0:])  // ABCDEF
	fmt.Println(abc[3:5]) // DE
	fmt.Println(abc[3])   // 68
}

func ContainString() {
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

	// Contain Rune
	// it reports whether a rune is within a string.
	result := strings.ContainsRune("golang", 'g')
	fmt.Println(result) // => true
}

func HasPrefix() {
	newString := "Go is a great programming language, and Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go"))    // true
	fmt.Println(strings.HasSuffix("Here we Go", "Go")) // true
	fmt.Println(strings.Count(newString, "Go"))        // 2
	fmt.Println(strings.LastIndex(newString, "Go"))    // 40
	fmt.Println(strings.Index(newString, "Go"))        // 0
}

func TrimString() {
	str1 := strings.TrimSpace("\t   Goodbye Windows, Welcome Linux!    ")
	fmt.Println(str1) // Goodbye Windows, Welcome Linux!

	str2 := strings.Trim("...Hello! Gos!!!!?", ".!?")
	fmt.Println(str2) // Hello! Gos
}
