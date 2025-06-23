package main

import "fmt"

const StatusOK = 200
const (
	StatusCreated  = 201
	StatusAccepted = 202
)

func main() {
	fmt.Println("Status OK:", StatusOK)
	fmt.Println("Status Created:", StatusCreated)
	fmt.Println("Status Accepted:", StatusAccepted)
}
