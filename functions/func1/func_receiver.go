package func1

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// receiver
func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs\n", a.Name, a.NumberOfLegs)
}

func addTwoNumbers(x, y int) int {
	return x + y
}

func addTwoNumberSum(x, y int) (sum int) {
	sum = x + y
	return
}

func sumMany(nums ...int) int {
	total := 0

	for _, i := range nums {
		total += i
	}

	return total
}

func ReceiverTest() {
	z := addTwoNumbers(2, 4)
	fmt.Println(z)

	x := addTwoNumberSum(5, 6)
	fmt.Println(x)

	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println(myTotal)

	// using Animal
	dog := Animal{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}
	dog.Says()

	var cat Animal
	cat.Name = "cat"
	cat.Sound = "meow"
	cat.NumberOfLegs = 4

	cat.Says()
	cat.HowManyLegs()
}
