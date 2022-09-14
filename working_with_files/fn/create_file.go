package fn

import (
	"fmt"
	"log"
	"os"
)

func CreateAFile() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("atest.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	} else {
		fmt.Println(newFile)
	}

	// empty a file
	err = os.Truncate("atest.txt", 0)
	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()
}
