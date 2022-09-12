package slices1

import (
	"fmt"
	"sort"

	"endangismaya.com/gotest/utils"
	"golang.org/x/exp/slices"
)

func AppendingIterateSortDeleteSlices() {

	utils.PrintTitle("Appending, Iterate, Sort, And Delete")

	// Appending
	var animals []string

	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals) // [dog fish cat horse]

	// Iterate
	for idx, item := range animals {
		fmt.Println(idx, item)
	}
	// out:
	// 0 dog
	// 1 fish
	// 2 cat
	// 3 horse

	// get single element
	fmt.Println("Element 0 is", animals[0])                     // out: Element 0 is dog
	fmt.Println("First two elements are", animals[0:2])         // out: First two elements are [dog fish]
	fmt.Println("The last element is", animals[len(animals)-1]) // out: The last element is horse

	// sorting slices
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))     // out: Is is sorted? false
	sort.Strings(animals)                                            // sorting a slice
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals)) // Is it sorted now? true
	fmt.Println(animals)                                             // out: [cat dog fish horse]

	// animals = DeleteFromSlice(animals, 1)
	// fmt.Println(animals)
	animals = slices.Delete(animals, 0, 1)
	fmt.Println(animals) // [dog fish horse]

}
