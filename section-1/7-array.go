package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)

	var gadget = [3]string{"mobile", "laptop", "tablet"}
	fmt.Println("Array 3:", gadget)
	fmt.Println("Panjang Array 3:", len(gadget))
	fmt.Println("Elemen pertama Array 3:", gadget[0])
}
