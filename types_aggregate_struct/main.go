package main

import (
	"endangismaya.com/gotest/types_aggregate_struct/struct1"
	"endangismaya.com/gotest/utils"
)

func main() {
	utils.PrintTitle("Car Implementation")
	struct1.CarImplementation()

	utils.PrintTitle("Book Implementation")
	struct1.BookImplementation()
	struct1.EmbeddedStruct()
}
