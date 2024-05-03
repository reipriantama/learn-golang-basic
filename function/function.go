package main

import "fmt"

func sayHello() string {
	name := "Rei"

	return name
}

func sayHelloTo(fristName string, lastName string) {
	fmt.Println("Hello", fristName, lastName)
}

func main() {
	sayHello()
	sayHelloTo("Rei", "Priantama")

}
