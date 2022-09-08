package main

import "fmt"

// reference type (no need pointers when working with maps)
// map

func main() {
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// iterate
	// for key, value := range intMap {
	// 	fmt.Println(key, value)
	// }

	// delete(intMap, "four")
	// for key, value := range intMap {
	// 	fmt.Println(key, value)
	// }

	// check if element exists in map
	element, ok := intMap["four"]
	if ok {
		fmt.Println(element, "is in map")
	}

	element1 := GetMap(intMap, "six", 6)
	fmt.Println(element1)

	// modify map
	intMap["five"] = 55
	fmt.Println(intMap)
}

func GetMap(map1 map[string]int, term string, default1 int) int {
	element, ok := map1[term]
	if ok {
		return element
	}

	return default1
}
