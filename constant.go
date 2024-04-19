package main

import "fmt"

func main() {
	// const firstName string = "Reinato"
	// const lastName = "Priantama"

	const (
		firstName string = "Rei"
		lastName         = "Priantama"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// ketika membuat constant dan tidak digunakan tidak akan error berbeda dengan variable
	// const tidak bisa mengubah data

	// error
	// fristName = "Rei"
	// lastName = "Prian"

	// * constant tidak bisa diubah variable bisa diubah
}
