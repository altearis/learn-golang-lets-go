package main

import "fmt"

func main() {
	person := map[string]any{
		"name":    "Indra",
		"address": "Jakarta",
	}
	fmt.Println("Nama:", person["name"])
	fmt.Println("Alamat :", person["address"])
	fmt.Println("No Hp :", person["hp"])
	fmt.Println("--------------------")
	person["hp"] = "081234567890"
	fmt.Println("No HP:", person["hp"])
	fmt.Println("--------------------")

	person2 := []map[string]any{
		{
			"name":    "Budi",
			"address": "Bandung",
		},
	}
	fmt.Println("Person 2 Name:", person2[0]["name"])
	person2 = append(person2, person)
	fmt.Println("Person 3 Name:", person2[1]["name"])
}
