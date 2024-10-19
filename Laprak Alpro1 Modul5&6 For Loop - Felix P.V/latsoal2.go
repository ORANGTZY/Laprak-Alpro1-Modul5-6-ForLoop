package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	// Menghitung volume untuk setiap kerucut
	for i := 1; i <= n; i++ {
		var jariJari, tinggi float64
		fmt.Printf("Masukkan jari-jari dan tinggi kerucut ke-%d: ", i)
		fmt.Scan(&jariJari, &tinggi)

		// Rumus volume kerucut: V = 1/3 * π * r^2 * t
		volume := (1.0 / 3.0) * math.Pi * math.Pow(jariJari, 2) * tinggi

		// Mencetak volume kerucut
		fmt.Printf("Volume kerucut ke-%d adalah: %.14f\n", i, volume)
	}
}
