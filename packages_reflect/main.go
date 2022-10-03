package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}

	return true
}

func main() {
	sample := Sample{"Endang"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())    // 1
	fmt.Println(sampleType.Field(0).Name) // Name

	fmt.Println(sampleType.Field(0).Tag.Get("required")) // true
	fmt.Println(sampleType.Field(0).Tag.Get("max"))      // 10

	fmt.Println(isValid(sample)) // true
}
