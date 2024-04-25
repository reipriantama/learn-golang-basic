package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("unknown type")
	}
}
