package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRei NoKTP = "111111111"

	var ktpAdam string = "327312232"
	var contoh string = "222222222"
	var contohKtp NoKTP = NoKTP(contoh)
	var contohKtpAdam NoKTP = NoKTP(ktpAdam)

	fmt.Println(ktpRei)
	fmt.Println(ktpAdam)
	fmt.Println(contohKtp)
	fmt.Println(contohKtpAdam)

}
