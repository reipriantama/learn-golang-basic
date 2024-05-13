package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2002, 4, 23, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-12-05 18:13:45"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())

	fmt.Println("start ")
	time.Sleep(4 * time.Second)
	fmt.Println("after 4 seconds")

	count := 0

	for {
		fmt.Println("Hello")
		count++

		if count == 5 {
			break
		}

		time.Sleep(2 * time.Second)
	}
}
