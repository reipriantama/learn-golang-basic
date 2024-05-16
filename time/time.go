package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now() // time.Now() untuk mendapatkan waktu saat ini
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2002, 4, 23, 0, 0, 0, 0, time.UTC) // time.Date() untuk membuat waktu
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05" // format penulisan waktu sesuai dengan RFC

	value := "2020-12-05 18:13:45"
	valueTime, err := time.Parse(formatter, value) // time.Parse() untuk mengkonversi string ke waktu
	if err == nil {
		fmt.Println(valueTime)
	} else {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())

}
