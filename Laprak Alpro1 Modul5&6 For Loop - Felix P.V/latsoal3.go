package main

import "fmt"

func main() {
	var base, exponent, result int
	fmt.Print("Masukkan bilangan dasar (base): ")
	fmt.Scan(&base)
	fmt.Print("Masukkan bilangan pangkat (exponent): ")
	fmt.Scan(&exponent)

	// Inisialisasi hasil dengan 1 (karena bilangan apapun dipangkatkan 0 hasilnya 1)
	result = 1

	// Menghitung base^exponent menggunakan perkalian berulang
	for i := 1; i <= exponent; i++ {
		result *= base
	}

	// Mencetak hasil pemangkatan
	fmt.Printf("Hasil dari %d pangkat %d adalah: %d\n", base, exponent, result)
}
