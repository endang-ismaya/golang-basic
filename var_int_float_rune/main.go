package main

import "fmt"

func main() {
	// int type
	var i1 int8 = 100 // -129 will be error
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	// float
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// byte alias for uint8
	// rune alias for int32

	// rune
	var r rune = 'f'
	fmt.Printf("%T\n", r) // : out int32
	fmt.Println(r)        // out:
	fmt.Printf("%x\n", r) // 66

	var idata rune = 10
	fmt.Printf("%T %d\n", idata, idata)

	// string
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)
}
