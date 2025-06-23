package main

import "fmt"

func calc(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
}
func calcV2(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}
func main() {
	s, d := calc(10, 20)
	fmt.Println("Jumlah : ", s)
	fmt.Println("Selisih :", d)

}
