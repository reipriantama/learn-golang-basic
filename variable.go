package main

import "fmt"

func main() {
	// var name string

	// name = "Rei"
	// fmt.Println(name)

	// ataupun bisa langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya

	// var name = "Reinato"
	// fmt.Println(name)

	// penulisan yang paling simple, yaitu dengan tanda :=
	name := "Reipriantama"
	fmt.Println(name)

	name = "Reinato Priantama"
	fmt.Println(name)

	var (
		firstName = "Rei"
		lastName  = "Priantama"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName, lastName)

}
