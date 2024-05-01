package main

import "fmt"

func thirdMax(nums []int) int {
    firstMax, secondMax, thirdMax := nums[0], nums[0], nums[0]
    count := 0

    // Temukan nilai maksimum pertama
    for _, num := range nums {
        if num > firstMax {
            firstMax = num
        }
    }

    // Temukan nilai maksimum kedua
    for _, num := range nums {
        if num < firstMax && num > secondMax {
            secondMax = num
        }
    }

    // Temukan nilai maksimum ketiga (jika ada)
    for _, num := range nums {
        if num < secondMax && num > thirdMax {
            thirdMax = num
            count++
        }
    }

    // Jika tidak ada nilai maksimum ketiga yang unik, kembalikan nilai maksimum pertama
    if count < 2 {
        return firstMax
    }

    return thirdMax
}

func main() {
    nums1 := []int{3, 2, 1}
    fmt.Println("Output 1:", thirdMax(nums1)) 

    nums2 := []int{1, 2}
    fmt.Println("Output 2:", thirdMax(nums2)) 

    nums3 := []int{2, 2, 3, 1}
    fmt.Println("Output 3:", thirdMax(nums3)) 
}
