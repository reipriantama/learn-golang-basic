package main

import "fmt"

func getFullName() (string, string) {
	return "Rei", "Priantama"
}

func getNameNumber() (string, int) {
	return "Reinato", 123123
}

func getNumber1(name string, lastName string, number int) (string, string, int) {
	return name, lastName, number

}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	getName, getNumber := getNameNumber()
	fmt.Println(getName, getNumber)

	// mengambil hanya satu return value
	firstName, _ = getFullName()
	fmt.Println(firstName)

	//-------------------------------//

	fmt.Println(getNumber1("Rei"+"nato", "priantama", 230499))
}
