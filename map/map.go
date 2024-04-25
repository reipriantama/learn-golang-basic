package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Rei",
		"address": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Rei"
	book["ups"] = "Salah"

	fmt.Println(book)

	fmt.Println(len(book))
	fmt.Println(book["title"])
	delete(book, "ups")
	fmt.Println(book)
}
