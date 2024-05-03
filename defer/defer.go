package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // akan dijalankan setelah function dijalankan(akan berada di akhir function)
	fmt.Println("Run Application")

}

func main() {
	runApplication()
}
