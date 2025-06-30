package main

import (
	"errors"
	"fmt"
)

func main() {
	tanpaPointer := 10
	var denganPointer *int = &tanpaPointer

	fmt.Println("Nilai x:", tanpaPointer)
	fmt.Println("Alamat x:", denganPointer)
	fmt.Println("Nilai dari pointer p:", *denganPointer)
	fmt.Println("--------------------")

	nilaiA := 200
	nilaiB := 300
	fmt.Println("Hasil Perhitungan:", hitungPenjumlahan(&nilaiA, &nilaiB))
	fmt.Println("--------------------")

	panjang := 10.0
	lebar := 5.0

	hasil, err := hitungLuasPersegiPanjang(panjang, lebar)

	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	fmt.Printf("Panjang: %.2f, Lebar: %.2f\n", panjang, lebar)
	fmt.Printf("Luas persegi panjang: %.2f\n", *hasil)
}

// Pointer sebagai parameter function. (Pass by Reference)
func hitungPenjumlahan(a *int, b *int) int {
	return *a + *b
}

// Pointer sebagai response data (Returning by Reference)
func hitungLuasPersegiPanjang(panjang, lebar float64) (*float64, error) {
	if panjang <= 0 || lebar <= 0 {
		return nil, errors.New("panjang dan lebar harus lebih dari 0")
	}

	luas := panjang * lebar
	return &luas, nil
}
