package main

import "fmt"

func reverse(n int) int {
	data := 0
	for n > 0 {
		remainder := n % 10
		data *= 10
		data += remainder
		n /= 10
	}
	return data
}

func main() {
	num := 123
	fmt.Println(reverse(num))
}
