package main

import (
	"fmt"
)

type Customer struct {
	Name, Adress string
	Age          int
}

type Student struct {
	Name  string
	Grade int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func (s Student) sayHelloStudent() {
	fmt.Println("Hello my name is", s.Name, "and my grade is", s.Grade)
}

func main() {
	var s1 Student
	s1.Name = "Rei"
	s1.Grade = 1

	fmt.Println("name :", s1.Name)
	fmt.Println("grade :", s1.Grade)

	var s2 = Student{"Ethan", 2}

	var s3 = Student{Name: "jason", Grade: 3}

	fmt.Println("student 1 :", s1.Name, "grade :", s1.Grade)
	fmt.Println("student 2 :", s2.Name, "grade :", s2.Grade)
	fmt.Println("student 3 :", s3.Name, "grade :", s3.Grade)

	var s4 = Student{Name: "Rei", Grade: 21}
	s4.sayHelloStudent()

	var rei Customer
	rei.Name = "Reinato Priantama"
	rei.Adress = "Bandung"
	rei.Age = 22

	fmt.Println(rei)
	fmt.Println(rei.Name)
	fmt.Println(rei.Adress)
	fmt.Println(rei.Age)

	rei.sayHello("Agus")

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
