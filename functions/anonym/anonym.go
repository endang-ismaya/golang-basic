package anonym

import "fmt"

func AnonymTest() {
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!")

	a := increment(10)
	fmt.Printf("%T\n", a)

	a()              // hos function is called
	fmt.Println(a()) // 12
}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
