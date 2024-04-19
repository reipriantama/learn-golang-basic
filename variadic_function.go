package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// jika menggunakan slice
func sumAllSlice(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := sumAll(1, 2, 3)
	fmt.Println(result)

	// secara langsung
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sumAll(100, 100, 100, 100, 100, 100))

	fmt.Println(sumAllSlice([]int{10, 10, 10, 10, 10, 10}))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

}
