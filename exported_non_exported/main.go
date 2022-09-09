package main

import (
	"fmt"

	"endangismaya.com/gotest/exported_non_exported/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allen", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	fmt.Println(myStaff.All())

	// staff.OverPaidLimit = 50000
	fmt.Println("Overpaid staff: ", myStaff.Overpaid())
	fmt.Println("Underpaid staff: ", myStaff.Underpaid())
}
