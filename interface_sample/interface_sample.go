package main

import "fmt"

type kalkulasiGaji interface {
	gaji() int
}

type pegawaiTetap struct {
	nama      string
	gajiPokok int
	tunjangan int
}

func (pt pegawaiTetap) gaji() int {
	return pt.gajiPokok + pt.tunjangan
}

type pegawaiKontrak struct {
	nama      string
	gajiPokok int
}

func (pk pegawaiKontrak) gaji() int {
	return pk.gajiPokok
}

type freelancer struct {
	nama       string
	ratePerJam int
	totalJam   int
}

func (fr freelancer) gaji() int {
	return fr.ratePerJam * fr.totalJam
}

func totalPengeluaran(kalk []kalkulasiGaji) int {
	total := 0

	for _, k := range kalk {
		total += k.gaji()
	}

	return total
}

func main() {
	pt1 := pegawaiTetap{"a", 5000000, 2000000}
	pt2 := pegawaiTetap{"b", 5000000, 3000000}

	pk1 := pegawaiKontrak{"c", 4000000}
	pk2 := pegawaiKontrak{"d", 4500000}

	fr1 := freelancer{"e", 35000, 80}
	fr2 := freelancer{"f", 35000, 70}

	fmt.Println(pt1.gaji())
	fmt.Println(pt2.gaji())

	fmt.Println(pk1.gaji())
	fmt.Println(pk2.gaji())

	fmt.Println(fr1.gaji())
	fmt.Println(fr2.gaji())

	totalPengeluaran := totalPengeluaran([]kalkulasiGaji{pt1, pt2, pk1, pk2, fr1, fr2})
	fmt.Println(totalPengeluaran)
}
