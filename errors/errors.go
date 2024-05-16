package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh dengan NOL")
	} else {
		return nilai / pembagi, nil //karena error sebuah interface, maka bisa dikembalikan dengan nil
	}

}

func main() {
	hasil, err := pembagian(100, 20)
	if err == nil {
		fmt.Println("Hasil", hasil)

	} else {
		fmt.Println("Error", err.Error())

	}
}
