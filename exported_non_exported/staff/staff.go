package staff

import "fmt"

// package level variable
var OverPaidLimit = 75_000  // exported with Uppercase
var underPaidLimit = 20_000 // not exported with lowercase

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}

	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary < underPaidLimit {
			underpaid = append(underpaid, x)
		}
	}

	e.notAPublic()
	notAReceiverFunction()
	return underpaid
}

func (e *Office) notAPublic() {
	fmt.Println("This is a private function")
}

func notAReceiverFunction() {
	fmt.Println("This is a not Receiver Function")
}
