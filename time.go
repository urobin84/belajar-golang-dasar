package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println("Tahun : ", now.Year())
	fmt.Println("Bulan : ", now.Month())
	fmt.Println("Hari : ", now.Day())
	fmt.Println("jam : ", now.Hour())
	fmt.Println("Menit : ", now.Minute())
	fmt.Println("Detik : ", now.Second())
	fmt.Println(now.Nanosecond())

	fmt.Println(now.UTC())
	fmt.Println(now.Zone())
}
