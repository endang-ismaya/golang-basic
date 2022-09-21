package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		log.Printf("%s is DOWN!\n", url)
	} else {
		defer res.Body.Close()
		log.Printf("%s -> Status Code: %d\n", url, res.StatusCode)

		if res.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			file := strings.Split(url, "//")[1] // http://www.google.com
			file += ".txt"

			log.Printf("writing response body to %s\n", file)

			// err = io.WriteFile(file, bodyBytes, 0664)
			err = os.WriteFile(file, bodyBytes, 0644)
			if err != nil {
				log.Fatal(err)
			}

		}
	}

	// Go routine completed
	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	// create WaitGroup
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}

	// wait for go routines
	wg.Wait()
}
