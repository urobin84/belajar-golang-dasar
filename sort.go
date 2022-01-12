package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Robin", 25},
		{"Joko", 35},
		{"Eko", 15},
		{"Budi", 20},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
