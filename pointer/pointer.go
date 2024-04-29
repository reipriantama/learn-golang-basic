package main

import "fmt"

type address struct {
	City, Province, Country string
}

func main() {
	// address1 := address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// -------- > Pointer

	address1 := address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //Pointer

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	var x int = 10

	// Mendeklarasikan pointer yang menunjuk ke variabel x
	var ptr *int = &x

	// Menampilkan nilai dari variabel x
	fmt.Println("Nilai dari x:", x)

	// Menampilkan alamat memori dari variabel x
	fmt.Println("Alamat memori dari x:", &x)

	// Menampilkan nilai yang ditunjuk oleh pointer
	fmt.Println("Nilai yang ditunjuk oleh pointer:", *ptr)

	// Mengubah nilai variabel melalui pointer
	*ptr = 20

	// Menampilkan nilai variabel setelah diubah melalui pointer
	fmt.Println("Nilai dari x setelah diubah melalui pointer:", x)
}
