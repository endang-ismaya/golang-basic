package main

import "fmt"

func main() {
	age := 10
	name := "jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t\n", name, age, rightHanded)

	ageInTenYears := age + 10 // age + 10 is expression because it can be valuated to a single value
	isATeenager := age >= 13

	fmt.Printf("In ten years, %s will be %d years old.\n", name, ageInTenYears)
	fmt.Println(name, "is a teenager:", isATeenager)

}
