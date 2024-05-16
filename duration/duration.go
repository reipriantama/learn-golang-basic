package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 60 * time.Second
	var duration2 time.Duration = 5 * time.Second
	var duration3 time.Duration = duration1 - duration2
	var duration4 time.Duration = duration1 + duration2

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3)
	fmt.Println(duration4)

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

		time.Sleep(1 * time.Second)
	}

}
