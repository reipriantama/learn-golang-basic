package main

import "fmt"

func Ups() any {
	// return 1
	// return true
	return "Rei"
}

func main() {

	var kosong any = Ups()
	fmt.Println(kosong)

	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = 1234
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)
}
