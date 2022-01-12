package main

import (
	"fmt"
	"strconv"
)

func main() {
	var boolean, err = strconv.ParseBool("true")

	if err == nil{
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(10000, 10)
	fmt.Println(value)

	valueInt, err := strconv.Atoi("300000")
	fmt.Println(valueInt)

}
