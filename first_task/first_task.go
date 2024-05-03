// Given an integer x, return true if x is a palindrome, and false otherwise.

package main

import "fmt"

func reverse(n int) int {
	data := 0
	for n > 0 {
		remainder := n % 10 //121 %10 = 1  // looping ulang menjadi 12
		data *= 10          // 0 * 10 = 0 // 0 * 10 = 0
		data += remainder   // 0 + 1 = 1 // 0 + 2 = 2
		n /= 10             // 121 / 10 = 12 // 12 / 10 = 1
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

	num1 := 121
	fmt.Println(isPalindrome(num1))
	fmt.Println(reverse(num1))

	// num2 := -121
	// fmt.Println(isPalindrome(num2))
	// fmt.Println(reverse(num2))

	// num3 := 123
	// fmt.Println(isPalindrome(num3))
	// fmt.Println(reverse(num3))

}
