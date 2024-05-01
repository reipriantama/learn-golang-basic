package main

import "fmt"

func validMountainArray(arr []int) bool {
    n := len(arr)
    if n < 3 {
        return false
    }

    // Cari puncak gunung
    peakIndex := 0
    for peakIndex < n-1 && arr[peakIndex] < arr[peakIndex+1] {
        peakIndex++
    }

    // Jika puncak berada di ujung larik atau di awal larik, bukan gunung
    if peakIndex == 0 || peakIndex == n-1 {
        return false
    }

    // Periksa bagian kanan dari puncak, harus menurun
    for i := peakIndex; i < n-1; i++ {
        if arr[i] <= arr[i+1] {
            return false
        }
    }

    return true
}

func main() {
    arr1 := []int{2, 1}
    fmt.Println("Output 1:", validMountainArray(arr1)) // Output: false

    arr2 := []int{1,2,3,3,4,5,6}
    fmt.Println("Output 2:", validMountainArray(arr2)) // Output: false

    arr3 := []int{0, 3, 2, 1}
    fmt.Println("Output 3:", validMountainArray(arr3)) // Output: true
}
