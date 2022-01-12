package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Muqorrobin", "robin"))
	fmt.Println(strings.Split("Muhammad Muqorrobin", " "))
	fmt.Println(strings.ToLower("Muhammad Muqorrobin"))
	fmt.Println(strings.ToUpper("Muhammad Muqorrobin"))
	fmt.Println(strings.Trim("    Muhammad Muqorrobin   ", " "))
	fmt.Println(strings.ReplaceAll("NO NO NO", "NO", "YES"))
}
