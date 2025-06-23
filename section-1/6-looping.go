package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Iterasi ke-", i)
	}
	fmt.Println("-------------------")

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Nama Buah : %s\n", index, fruit)
	}
	fmt.Println("-------------------")

	// Range with map
	scores := map[string]int{"John": 90, "Alice": 85, "Bob": 78}
	for name, score := range scores {
		fmt.Printf("Name: %s, Score: %d\n", name, score)
	}
}
