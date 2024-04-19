package main

import "fmt"

// func sayHellowWithFilter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello", filteredName)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// ------------------------ penulisan lebih effisien

type Filter func(string) string

func sayHellowWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHellowWithFilter("Rei", spamFilter)

	filter := spamFilter
	sayHellowWithFilter("Anjing", filter)
}
