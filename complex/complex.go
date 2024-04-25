package main

import "fmt"

func main() {
	var a complex64 = 5 + 12i
	var b complex64 = 2 + 3i

	// Penambahan
	fmt.Println("Penambahan:", a+b)

	// Pengurangan
	fmt.Println("Pengurangan:", a-b)

	// Perkalian
	fmt.Println("Perkalian:", a*b)

	// Pembagian
	fmt.Println("Pembagian:", a/b)

	var c complex128 = 5 + 12i
	var d complex128 = 2 + 3i

	// Penambahan
	fmt.Println("Penambahan:", c+d)

	// Pengurangan
	fmt.Println("Pengurangan:", c-d)

	// Perkalian
	fmt.Println("Perkalian:", c*d)

	// Pembagian
	fmt.Println("Pembagian:", c/d)
}
