package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rei := Man{"Rei"}
	rei.Married()

	fmt.Println(rei.Name)
}
