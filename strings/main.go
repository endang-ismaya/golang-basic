package main

import (
	"fmt"
	"strings"

	"endangismaya.com/gotest/strings/manipulate"
)

func main() {
	// Concatenate String
	manipulate.ConcatenateString()

	// Getting Substring
	manipulate.GetSubstring()

	// Contains
	manipulate.ContainString()

	// HasPrefix
	manipulate.HasPrefix()

	// repeat and replace
	manipulate.RepeatAndReplaceString()

	// compare string insensitive case
	fmt.Println(strings.EqualFold("GO", "go")) // true

	// Split String
	manipulate.SplitString()

	// Join Slice
	manipulate.JoinSlices()

	// Fields
	manipulate.FieldsString()

	// Trim String
	manipulate.TrimString()

}
