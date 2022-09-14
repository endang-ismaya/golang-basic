package main

import (
	"endangismaya.com/gotest/utils"
	"endangismaya.com/gotest/working_with_files/fn"
)

func main() {
	utils.PrintTitle("Creating a File")
	fn.CreateAFile()

	utils.PrintTitle("Open A File")
	fn.OpenAFile()
	fn.OpenFileMode()

	utils.PrintTitle("Get File Info")
	fn.GetFileInfo()

	utils.PrintTitle("Read file line by line")
	fn.ReadFileLineByLine()
}
