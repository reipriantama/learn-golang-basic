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

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverse(x)
}

func main() {
	num := 123456
	fmt.Println(num, isPalindrome(num))
	fmt.Println(reverse(num))

}
