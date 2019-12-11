package main

import (
	"fmt"
)

func stopOn5(f func(x int) int) func(y int) {
	x := 0
	return func(y int) {
		if x >= 5 {
			fmt.Println("Function already ran five times")
		} else {
			x++
			fmt.Println(f(y))
		}

	}

}
func increment(x int) int {
	return x + 1
}

func main() {
	wrappedFunc := stopOn5(increment)
	wrappedFunc(10)
	wrappedFunc(10)
	wrappedFunc(10)
	wrappedFunc(10)
	wrappedFunc(10)
	wrappedFunc(10)

}
