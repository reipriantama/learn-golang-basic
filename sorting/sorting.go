package main

import (
	"fmt"
	"sort"
)

func main() {
	number := []int{10, 8, 5, 6, 7, 2, 4, 1, 3, 9}

	sort.Ints(number)
	fmt.Println(number)

	var n = []int{1, 39, 2, 9, 7, 54, 11, 12}

	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				isDone = false
			}
			i++
		}
	}

	fmt.Print(n)

}
