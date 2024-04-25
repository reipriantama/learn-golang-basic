package main

import "fmt"

func main() {
	var name = "Rei"
	var age = 22
	var number float32 = 3.14

	fmt.Println(name)

	fmt.Println("Hello", name, "your age is", age, "and your number is", number)

	fmt.Printf("Type Data: %T \n", name)
	fmt.Printf("Value Data: %v \n", name)
	fmt.Printf("Hello %v your age is %v and your number is %v \n", name, age, number)

	fmt.Printf("Type Data: %T dan value: %v \n", number, number)
}
