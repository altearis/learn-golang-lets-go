package main

import "fmt"

func main() {
	firstSlice := []int{1}
	fmt.Printf("Current firstSlice : %v", firstSlice)
	fmt.Printf("\nLength of firstSlice : %d\n", len(firstSlice))
	fmt.Println("--------------------")
	secondSlice := []int{2, 3, 4}
	fmt.Printf("Current secondSlice : %v", secondSlice)
	fmt.Printf("\nLength of secondSlice : %d\n", len(secondSlice))
	fmt.Println("--------------------")
	thirdSlice := append(firstSlice, secondSlice...)
	fmt.Printf("Current thirdSlice : %v", thirdSlice)
	fmt.Printf("\nLength of thirdSlice : %d\n", len(thirdSlice))
	fmt.Println("--------------------")
}
