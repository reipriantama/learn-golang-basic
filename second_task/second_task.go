/* Given an integer array nums, return the third distinct maximum number in
this array. If the third maximum does not exist, return the maximum number. */

package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	firstMax, secondMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num > firstMax {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = num
		} else if num < firstMax && num > secondMax {
			thirdMax = secondMax
			secondMax = num
		} else if num < secondMax && num > thirdMax {
			thirdMax = num
		}
	}

	if thirdMax != math.MinInt64 {
		return thirdMax
	}
	return firstMax
}

func main() {

	nums3 := []int{9, 5, 7, 4, 2}
	fmt.Println("Output 3:", thirdMax(nums3))
}
