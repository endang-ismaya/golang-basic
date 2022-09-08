package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public (or exported)" // with capital letter

func notExported() {
	fmt.Println("Not Exported")
}

func Exported() {
	fmt.Println("Inside Exported:")
	fmt.Println("Exported")
	fmt.Println(privateVar)

	notExported()
}
