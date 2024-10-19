package main

import "fmt"

func main() {
	var n, total int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	// Menjumlahkan bilangan dari 1 hingga n
	for i := 1; i <= n; i++ {
		total += i
	}

	// Mencetak hasil penjumlahan
	fmt.Println("Hasil penjumlahan dari 1 sampai", n, "adalah", total)
}
