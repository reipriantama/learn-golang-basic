package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan sederhana ke", counter)
	}

	fmt.Println("Selesai")

	// loop
	names := []string{"Rei", "Eko", "Kurniawan"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Contoh loop", names[i])
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// loop range
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// loop range tanpa index
	for _, name := range names {
		fmt.Println("=", name)
	}
}
