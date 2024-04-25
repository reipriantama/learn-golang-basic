package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsesnsi = absensi > 80

	var lulus = lulusNilaiAkhir && lulusAbsesnsi
	fmt.Println(lulus)

}
