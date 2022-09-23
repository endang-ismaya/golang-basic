package main

import (
	"fmt"
	"strings"

	"endangismaya.com/gotest/strings/manipulate"
	"endangismaya.com/gotest/utils"
)

func main() {
	// Concatenate String
	utils.PrintTitle("Concatenate")
	manipulate.ConcatenateString()

	// Getting Substring
	utils.PrintTitle("Substring")
	manipulate.GetSubstring()

	// Contains
	utils.PrintTitle("Contains")
	manipulate.ContainString()

	// HasPrefix
	utils.PrintTitle("HasPrefix")
	manipulate.HasPrefix()

	// repeat and replace
	utils.PrintTitle("Repeat and Replace")
	manipulate.RepeatAndReplaceString()

	// compare string insensitive case
	utils.PrintTitle("Compare")
	fmt.Println(strings.EqualFold("GO", "go")) // true

	// Split String
	utils.PrintTitle("Split")
	manipulate.SplitString()

	// Join Slice
	utils.PrintTitle("Join")
	manipulate.JoinSlices()

	// Fields
	utils.PrintTitle("Fields")
	manipulate.FieldsString()

	// Trim String
	utils.PrintTitle("Trim")
	manipulate.TrimString()

}
