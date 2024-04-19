package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Rei"
	names[1] = "Reinato"
	names[2] = "Reinato Priantama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// names[3] = "Budi"
	// akan error karena ada penambahan index karena di awal sudah di set hanya bisa 3 index

	var values = [...]int{
		90,
		85,
		95,
		100,
		110,
		123,
	}

	fmt.Println(values)
	// bisa dipanggil satu per satu
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)

}
