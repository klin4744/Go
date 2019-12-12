package main

import (
	"fmt"
	"sync"
)

func sideGoRoutine(x int, wg *sync.WaitGroup) {
	fmt.Println("Go routine:", x)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go sideGoRoutine(0, &wg)
	go sideGoRoutine(1, &wg)

	fmt.Println("Main Goroutine")
	wg.Wait()
}
