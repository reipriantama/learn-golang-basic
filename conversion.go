package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Rei Priantama"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

	// karena melibihi kapasitas int16 maka nilai tersebut menjadi ke Nilai minimun int16 yaitu -32768
	// Tipe data int16 nilai minimun adalah -32768 nilai maksimum adalah 32767
}
