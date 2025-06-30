package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func taskA(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Task A")
	time.Sleep(2 * time.Second) // simulasi kerja Task A
	fmt.Println("Selesai Task A")
}

func taskB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Task B")
	time.Sleep(3 * time.Second) // simulasi kerja Task B
	fmt.Println("Selesai Task B")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)

	go taskA(&wg)
	go taskB(&wg)

	wg.Wait()
	duration := time.Since(start) // menghitung durasi eksekusi
	fmt.Println("Total waktu (dengan concurrency):", duration)
}
