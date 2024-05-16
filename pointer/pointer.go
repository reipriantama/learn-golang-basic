package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// -------- > Pointer

	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := &address1 //Pointer

	// address2.City = "Bandung"
	// fmt.Println(address1)
	// fmt.Println(address2)

	var x int = 10
	var y int = 20

	// Mendeklarasikan pointer yang menunjuk ke variabel x & y
	fmt.Println(x)
	x = 20
	fmt.Println(x)

	var ptr *int = &x
	var ptr2 *int = &y

	// Menampilkan nilai dari variabel x & y
	fmt.Println("Nilai dari x:", x)
	fmt.Println("Nilai dari x:", y)

	// Menampilkan alamat memori dari variabel x
	fmt.Println("Alamat memori dari x:", &x)
	fmt.Println("Alamat memori dari x:", &y)

	// Menampilkan nilai yang ditunjuk oleh pointer
	fmt.Println("Nilai yang ditunjuk oleh pointer:", *ptr)
	fmt.Println("Nilai yang ditunjuk oleh pointer:", *ptr2)

	// Mengubah nilai variabel melalui pointer
	*ptr = 20
	*ptr2 = 30

	// Menampilkan nilai variabel setelah diubah melalui pointer
	fmt.Println("Nilai dari x setelah diubah melalui pointer:", x)
	fmt.Println("Alamat dari x setelah diubah melalui pointer:", &x)

	fmt.Println("Nilai dari y setelah diubah melalui pointer:", y)
	fmt.Println("Alamat dari y setelah diubah melalui pointer:", &y)

}
