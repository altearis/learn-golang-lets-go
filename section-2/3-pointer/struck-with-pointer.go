package main

import "fmt"

type Student struct {
	Name  string
	Age   *int
	Major string
}

func main() {
	studentA := Student{
		Name:  "Indrabayu",
		Major: "Teknik Informatika",
	}

	fmt.Println(studentA)

	studentAage := 30
	studentA.Age = &studentAage
	fmt.Println(*studentA.Age)
}
