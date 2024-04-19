package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Rei"
	middleName = "Nato"
	lastName = "Priantama"

	return firstName, middleName, lastName
}

func getFullNameNumber() (firstName, lastName string, firstNumber int, secondNumber int) {
	firstName = "Rei"
	lastName = "Priantama"
	firstNumber = 123123
	secondNumber = 321321
	return firstName, lastName, firstNumber, secondNumber
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	d, e, f, _ := getFullNameNumber()
	fmt.Println(d, e, f)

}
