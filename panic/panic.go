package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	/* jika memanggil recover setelah error, recover akan berhenti
	karena tidak akan tereksekusi, harus berada di dalam defer
	*/
}

func main() {
	runApp(true)
	fmt.Println("Aplikasi berjalan setelah terjadi panic")
}
