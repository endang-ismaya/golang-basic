package manipulate

import (
	"fmt"
	"strings"
)

func ConcatenateString() {
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
}
