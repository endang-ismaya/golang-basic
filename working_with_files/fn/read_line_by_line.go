package fn

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFileLineByLine() {
	file, err := os.Open("E31AX_Parsed.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	success := scanner.Scan()

	if !success {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reach EOF")
		} else {
			log.Fatal(err)
		}
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal()
	}

}
