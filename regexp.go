package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("r([a-z])")

	fmt.Println(regex.MatchString("Budi"))
	fmt.Println(regex.MatchString("robin"))
	fmt.Println(regex.MatchString("Joko"))

	search := regex.FindAllString("joko budi robin", -1)
	fmt.Println(search)
}
