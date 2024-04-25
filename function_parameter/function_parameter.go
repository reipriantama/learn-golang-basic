package main

import "fmt"

func sayHelloTo(firstName string, lastName string, number int) {
	fmt.Println("Say hello to", firstName, lastName, number)
}

func main() {
	sayHelloTo("Rei", "Priantama", 230499)
}
