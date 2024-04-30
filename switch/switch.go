package main

import "fmt"

func main() {
	name := "Reinato"

	switch name {
	case "Rei":
		fmt.Println("Hello Rei")
	case "Eko":
		fmt.Println("Hello Eko")
	default:
		fmt.Println("Hello All")
	}

	// short statement switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama sudah benar")
	case false:
		fmt.Println("Nama tidak sesuai")
	}

	// switch without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama Sudah Benar")
	default:
		fmt.Println("Nama tidak sesuai")
	}
}
