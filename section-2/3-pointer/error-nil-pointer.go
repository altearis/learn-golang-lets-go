package main

import (
	"errors"
	"fmt"
)

type User struct {
	Nama string
}

func cariUser(id int) (*User, error) {
	if id != 1 {
		return nil, errors.New("user tidak ditemukan")
	}

	return &User{Nama: "Faris"}, nil
}

func main() {
	user, err := cariUser(2) // user dengan ID 2 tidak ada
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// YANG BENAR
	// check variabel user!
	if user == nil {
		fmt.Println("User tidak ditemukan!")
		return
	}

	// ❌ Tetap coba akses user tanpa cek nil❌
	fmt.Println("Nama user:", user.Nama) // panic: nil pointer dereference
}
