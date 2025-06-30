package main

import "fmt"

// Class/Object Account
type Account struct {
	nama     string // private
	Username string // public
}

// encapsulated set access
func (u *Account) SetNama(n string) {
	u.nama = n
}

// encapsulated read access
func (u Account) GetNama() string {
	return u.nama
}

func main() {
	u := Account{Username: "ibp"}
	u.SetNama("Indrabayu aja")

	fmt.Println("Username:", u.Username)
	fmt.Println("Nama lengkap:", u.GetNama())
	fmt.Println(u.nama)
}
