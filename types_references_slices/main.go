package main

import (
	slices1 "endangismaya.com/gotest/types_references_slices/slices_1"
)

// reference type (pointers, slices, maps, functions, channels)
// slices

func main() {

	// Appending, Iterate, Sort, and Delete Slice
	slices1.AppendingIterateSortDeleteSlices()

	// Comparing Slices
	slices1.ComparingAppendingSlices()

	// Copying Slices
	slices1.CopySlicesWithoutShared()

}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] // copy the last element to index i
	a[len(a)-1] = ""   // give empty value
	a = a[:len(a)-1]   // truncate last
	return a
}
