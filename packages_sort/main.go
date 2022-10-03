package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {

	users := UserSlice{
		{"Endang", 38},
		{"Indah", 37},
		{"Alde", 13},
		{"Aqeela", 10},
		{"Auza", 4},
		{"Arsyila", 2},
	}

	// before sort
	fmt.Println(users)

	sort.Sort(users)
	fmt.Println(users)
}
