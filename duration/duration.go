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

	start := time.Now()

	time.Sleep(3 * time.Second)

	duration := time.Since(start)

	fmt.Println("Waktu akan muncul dalam hitungan detik:", duration.Seconds())

}
