package main

import (
	"fmt"
	"sort"
)

func main() {
	number := []int{10, 8, 5, 6, 7, 2, 4, 1, 3, 9}

	sort.Ints(number)
	fmt.Println(number)
}
