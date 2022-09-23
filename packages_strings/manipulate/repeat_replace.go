package manipulate

import (
	"fmt"
	"strings"
)

func RepeatAndReplaceString() {
	// Repeat
	str1 := strings.Repeat("Ja", 10)
	fmt.Println(str1) // JaJaJaJaJaJaJaJaJaJa

	// Replace
	str2 := strings.Replace("192.168.0.1", ".", ":", 2)
	fmt.Println(str2) // 192:168:0.1

	str3 := strings.Replace("192.168.0.1", ".", ":", -1)
	fmt.Println(str3) // 192:168:0:1

	// ReplaceAll
	str4 := strings.ReplaceAll("192.168.0.1", ".", ":")
	fmt.Println(str4) // 192:168:0:1

}

func SplitString() {
	str5 := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", str5) // []string
}

func JoinSlices() {
	slice1 := []string{"I", "Love", "Golang"}
	joinSlice1 := strings.Join(slice1, "__")
	fmt.Println(joinSlice1) // I__Love__Golang
}

func FieldsString() {
	str6 := "Orang Blue Green \n Blue Yellow \r\n Black"
	fields := strings.Fields(str6)
	fmt.Printf("%T\n", fields)  // []string
	fmt.Printf("%#v\n", fields) // []string{"Orang", "Blue", "Green", "Blue", "Yellow", "Black"}
}
