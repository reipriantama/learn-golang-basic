package main

import "fmt"

func main() {
	var nilai8 int8 = 127
	var nilai16 int16 = 32767
	var nilai32 int32 = 2147483647
	var nilai64 int64 = 9223372036854775807

	var nilaiLebih = 128
	var nilaiBerubah int8 = int8(nilaiLebih)

	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilaiBerubah)

}
