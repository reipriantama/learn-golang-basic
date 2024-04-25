package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you're blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Rei", blacklist)

	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
