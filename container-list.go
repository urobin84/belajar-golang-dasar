package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Robin")
	data.PushBack("Yadi")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println("Element : ", element.Value)
	}

}
