package main

import "fmt"

func main() {
	// var x = 10

	var age int = 38
	fmt.Println("Age:", age)

	var name = "Endang"
	fmt.Println("My name is: ", name)
	fmt.Println(name[0]) // 69
	fmt.Println(name[1]) // 110

	// short var
	s := "Learning Golang"
	fmt.Println(s)

	car, cost := "Audi", 50_000
	fmt.Println(car, cost)

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var i, j int
	i, j = 5, 8
	fmt.Println(i, j)
}
