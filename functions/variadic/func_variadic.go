package variadic

import (
	"fmt"
	"strings"
)

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString
}

func VariadicFuncTest() {
	f1(1, 2, 3, 4)
	// []]int
	// []int{1, 2, 3, 4}
	f1()
	// []int
	// []int(nil)

	// passing a slice into variadic function
	nums := []int{1, 2, 3}
	f1(nums...)
	// []int
	// []int{1, 2, 3}
	f2(nums...)
	fmt.Println(nums) // [50 2 3]

	// mixing
	info := personInformation(35, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // => Age: 35, Full Name:Wolfgang Amadeus Mozart
}
