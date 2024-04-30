package main

import (
	"fmt"
	"learn-golang-basic/test"
)

func main() {
	result := test.TestFunction("Rei")
	fmt.Println(result)

	fmt.Println("John Wick")
	fmt.Println("John", "Wick")

	fmt.Print("John Wick \n")
	fmt.Print("John", "Wick \n")
}
