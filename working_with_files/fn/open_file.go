package fn

import (
	"fmt"
	"log"
	"os"
)

func OpenAFile() {
	file, err := os.Open("atest.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func OpenFileMode() {
	file, err := os.OpenFile("atest.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func GetFileInfo() {
	var fileInfo os.FileInfo

	fileInfo, err := os.Stat("atest.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exists!")
		}
	}
	fmt.Println("Filename:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Last modified:", fileInfo.ModTime().Local())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Permissions:", fileInfo.Mode())
}
