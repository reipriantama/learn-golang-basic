package main

import "fmt"

func main() {
	names := [...]string{"Rei", "Eko", "Kurniawan", "Kari", "Budi", "Nugraha"}

	slice1 := names[2:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"} // membuat array
	fmt.Println(days)

	daySlice1 := days[5:]
	fmt.Println("slice days :", daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println("mengganti data daySlice1 :", daySlice1)
	fmt.Println("data days setelah diubah :", days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println("daySlice2 :", daySlice2)

}
