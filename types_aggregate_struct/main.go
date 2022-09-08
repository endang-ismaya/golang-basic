package main

import "fmt"

// aggregate types (array, struct)

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	var myCar Car

	myCar.NumberOfTires = 4
	myCar.Luxury = false
	myCar.Make = "Volkswagen"
	myCar.Model = "Model X"

	myNewCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC",
		Year:          2019,
	}

	fmt.Println("My car is", myCar)
	fmt.Println("My new car is", myNewCar)
}
