package main

import (
	"fmt"
	"learn-golang-basic/test"
)

func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	result := test.TestFunction("Rei")
	fmt.Println(result)

	fmt.Println(sayHello("Rei"))

}

// unexported => hanya bisa digunakan pada package yang sama
// exported => bisa digunakan di package yang berbeda
