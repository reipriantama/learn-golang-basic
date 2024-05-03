package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter2 := 1; counter <= 20; counter++ {
		fmt.Println("Perulangan sederhana ke", counter2)
	}

	fmt.Println(counter)

	// loop
	names := []string{"Rei", "Eko", "Kurniawan"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Contoh loop :", names[i])
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

}
