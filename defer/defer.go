package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}
func logging2() {
	fmt.Println("Selesai memanggil function2")
}

func runApplication() {
	defer logging2()
	fmt.Println("Run Application")
	defer logging() // akan dijalankan setelah function dijalankan(akan berada di akhir function)
}

func main() {
	runApplication()
}
