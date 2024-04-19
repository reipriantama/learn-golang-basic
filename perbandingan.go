package main

import "fmt"

func main() {

	var name1 = "Rei"
	var name2 = "Rei"

	var result = name1 != name2
	fmt.Println(result)

	// jika tanda > bisa digunakan pada alphabet karena dengan urutan a,b,c....
	// tetapi hal tersebut jarangan digunakan

	var alphabet1 = "b"
	var alphabet2 = "a"
	var result1 = alphabet1 > alphabet2
	fmt.Println(result1)
}
