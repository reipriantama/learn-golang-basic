package main

import "fmt"

func main() {
	name := "Reinato"

	if name == "Rei" {
		fmt.Println("Hello Rei")
	} else if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hello All")
	}

	// short statement if
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah Benar")
	}
}
