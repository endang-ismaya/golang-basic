package utils

import (
	"fmt"
	"strings"
)

func PrintTitle(title string) {
	lineTitle := strings.Repeat("=", len(title)+1)
	fmt.Printf("\n# %s\n", lineTitle)
	fmt.Printf("# %s\n", title)
	fmt.Printf("# %s\n", lineTitle)
}
