package main

import (
	"endangismaya.com/gotest/functions/func1"
	"endangismaya.com/gotest/functions/variadic"
	"endangismaya.com/gotest/utils"
)

func main() {
	utils.PrintTitle("Func Basic")
	func1.FuncBasic()

	utils.PrintTitle("Func Receiver")
	func1.ReceiverTest()

	utils.PrintTitle("Func Variadic")
	variadic.VariadicFuncTest()
}
