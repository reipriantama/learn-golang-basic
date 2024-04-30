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
	sayHelloTo("Rei", "Priantama")

	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan", sayHello(), i)
	}
}
