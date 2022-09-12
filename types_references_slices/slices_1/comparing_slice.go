package slices1

import (
	"fmt"

	"endangismaya.com/gotest/utils"
	"golang.org/x/exp/slices"
)

func ComparingAppendingSlices() {

	utils.PrintTitle("Comparing Slice")

	// compare slices
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	c := []string{"a", "b", "c"}
	d := []string{"d", "e", "f"}
	fmt.Println("compare a & b", slices.Equal(a, b)) // compare a & b true
	fmt.Println("compare c & d", slices.Equal(c, d)) // compare c & d false

	// appending slices
	a = append(a, 4)
	a = append(a, 5, 6, 7, 8)
	fmt.Println(a) // [1 2 3 4 5 6 7 8]

	a = append(a, b...)
	fmt.Println(a) // [1 2 3 4 5 6 7 8 1 2 3]
}
