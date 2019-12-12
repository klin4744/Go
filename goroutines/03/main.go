// Using goroutines, create an incrementer program
// have a variable to hold the incrementer value
// launch a bunch of goroutines
// each goroutine should
// read the incrementer value
// store it in a new variable
// yield the processor with runtime.Gosched()
// increment the new variable
// write the value in the new variable back to the incrementer variable
// use waitgroups to wait for all of your goroutines to finish
// the above will create a race condition.
// Prove that it is a race condition by using the -race fla g

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Number of CPUs:\t", runtime.NumCPU())
	fmt.Println("Number of GoRoutines:\t", runtime.NumGoroutine())
	value := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			newValue := value
			runtime.Gosched()
			newValue++
			value = newValue
			wg.Done()

		}()
	}
	fmt.Println("Value:\t", value)
	fmt.Println("Number of CPUs:\t", runtime.NumCPU())
	fmt.Println("Number of GoRoutines:\t", runtime.NumGoroutine())

	wg.Wait()
}
