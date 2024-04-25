package main

import "fmt"

func main() {
	var nilaiUint8 uint8 = 255
	var nilaiUint16 uint16 = 65535
	var nilaiUint32 uint32 = 4294967295
	var nilaiUint64 uint64 = 18446744073709551615

	var nilaiLebih = 257
	var nilaiBerubah = uint8(nilaiLebih)

	fmt.Println(nilaiUint8)
	fmt.Println(nilaiUint16)
	fmt.Println(nilaiUint32)
	fmt.Println(nilaiUint64)
	fmt.Println(nilaiBerubah)
}
