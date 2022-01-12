package main

import "fmt"

func main()  {
	angka := []int{5, 4, 8, 9, 7, 3, 2}
	for i := 0; i < len(angka); i++ {
		fmt.Println(angka[i])
	}
}
