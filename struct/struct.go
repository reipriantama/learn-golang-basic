package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var rei Customer
	rei.Name = "Reinato Priantama"
	rei.Adress = "Bandung"
	rei.Age = 22

	fmt.Println(rei)
	fmt.Println(rei.Name)
	fmt.Println(rei.Adress)
	fmt.Println(rei.Age)

	joko := Customer{
		Name:   "Joko",
		Adress: "Jakarta",
		Age:    25,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Bandung", 23}

	fmt.Println(budi)

	budi.sayHello("Agus")
	joko.sayHello("Agus")
	rei.sayHello("Agus")
}
