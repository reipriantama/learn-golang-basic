package main

import "fmt"

func getFullNameAndAge() (string, int) {
	// Predefined value untuk nama lengkap
	defaultFullName := "Joko Budi"
	// Predefined value untuk umur
	defaultAge := 30

	// Mengembalikan nilai predefined dan nilai lainnya
	return defaultFullName, defaultAge
}

func main() {
	fullName, age := getFullNameAndAge()
	fmt.Println("Nama Lengkap:", fullName)
	fmt.Println("Umur:", age)
	fmt.Printf("Nama Lengkap: %v dan Umur: %v \n", fullName, age)
}
