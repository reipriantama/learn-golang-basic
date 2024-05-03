// Given an array of integers arr, return true if and only if it is a valid mountain array.

package main

import "fmt"

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	peakIndex := 0
	for peakIndex < n-1 && arr[peakIndex] < arr[peakIndex+1] {
		peakIndex++
	}

	if peakIndex == n-1 {
		return false
	}

	for i := peakIndex; i < n-1; i++ {
		if arr[i] <= arr[i+1] { // i = 1, arr[1] = 3, arr
			return false
		}
	}

	return true
}

func main() {

	arr3 := []int{0, 3, 2, 1, 0}
	fmt.Println("Output 3:", validMountainArray(arr3))
}
