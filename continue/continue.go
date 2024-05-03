package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke-", i)
	}

	names := [...]string{"Rei", "Eko", "Kurniawan", "Kari", "Budi", "Nugraha"}

	for i := 0; i < len(names); i++ {
		if i == 2 || i == 4 || i == 5 {
			continue

		}
		fmt.Println(names[i])
	}
}
