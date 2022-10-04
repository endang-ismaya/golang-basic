package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z]+)g")

	fmt.Println(regex.MatchString("Endang"))
	fmt.Println(regex.MatchString("endang"))
	fmt.Println(regex.MatchString("empang"))

	search := regex.FindAllString("endang enjang ekang wijaya ismaya aldeind", -1)
	fmt.Println(search)

	res := regex.FindAllStringSubmatch("endang is the student of Binus Learning Online", -1)
	fmt.Println(res)
	fmt.Println(res[0][1])

}
