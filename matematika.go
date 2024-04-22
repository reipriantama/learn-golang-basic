package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e
	fmt.Println(c)

	var i = 10

	i += 10 // sama aja dengan i = i + 10
	fmt.Println(i)

	i += 5 // sama aja dengan i = i + 5
	fmt.Println(i)

	c += 20
	fmt.Println(c)

	// Unary Operator
	// ++  ---> a = a + 1
	// -- ---> a = a - 1
	// - ---> Negative
	// + ---> Positive
	// ! ---> Boolean kebalikan

	var j = 1
	fmt.Println(j)

	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)

}
