package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name

}

func getNumber(number int) (string, int) {

	return "Hello ", number
}

func main() {
	contoh := getGoodBye

	fmt.Println(contoh("Rei"))
	fmt.Println(getNumber(123))

}
