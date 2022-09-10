package main

import "fmt"

func main() {
	answer := 7 + 3*4
	fmt.Println("Answer:", answer) // Answer: 19

	// precedence
	// multiplication and division
	a := 12.0 * 3.0 / 4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0 / 4.0)

	fmt.Println("a", a, "b", b, "c", c) // a 9 b 9 c 9

	// integer division
	unclear := 12 * (3 / 4)
	fmt.Println("unclear", unclear) // unclear 0

	// parentehesis
	f := 12.0 / 3.0 / 4.0 // 1
	fmt.Println(f)
	f = 12.0 / (3.0 / 4.0) // 16
	fmt.Println(f)
}
