package main

import "fmt"

func main() {
	var n, result int
	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	// Inisialisasi hasil faktorial dengan 1 (karena faktorial 0 = 1)
	result = 1

	// Menghitung faktorial menggunakan perulangan
	for i := 1; i <= n; i++ {
		result *= i
	}

	// Mencetak hasil faktorial
	fmt.Printf("Faktorial dari %d adalah: %d\n", n, result)
}
