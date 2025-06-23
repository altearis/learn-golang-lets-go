package main

import "fmt"

func main() {
	age := 20
	if age < 17 {
		fmt.Println("Belum cukup umur")

	} else if age == 17 {
		fmt.Println("Umur 17 tahun, sudah boleh ikut ujian SIM")
	} else {
		fmt.Println("Sudah ketuaan")
	}
	fmt.Println("-----------------------------")
	day := "Minggu"
	switch day {
	case "Senin", "Jumat":
		fmt.Println("Hari ini adalah Hari macet se Jakarta")
	case "Sabtu", "Minggu":
		fmt.Println("Hari ini adalah Hari libur....")
	default:
		fmt.Println("Hari ini adalah Hari biasa")
	}
	fmt.Println("-----------------------------")
	today := 20
	switch {
	case today <= 15:
		fmt.Println("Ngopi di Kopi Kenangan...")
	case today > 25:
		fmt.Println("Ngopi di Starbuck...")
	default:
		fmt.Println("Ngopi di rumah pake ampas kopi kemaren...")
	}
}
