package main

import "fmt"

func getName(name string) string {
	hello := "Hello " + name
	return hello
}

func getNumber(number int) int {
	return number
}

func main() {
	result := getName("Rei")
	fmt.Println(result)

	resultNumber := getNumber(100)
	fmt.Println(resultNumber)

	// atau lakukan secara langsung
	fmt.Println(getName("Budi"))
	fmt.Println(getName("Joko"))
	// --------------------------------- //
	fmt.Println(result, resultNumber)
	fmt.Println(getName("Dimas"), getNumber(2000))
}
