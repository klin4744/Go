package main

import "fmt"

func main() {
	c := make(chan int)
	populateChannel((chan<- int)(c))
	pullFromChan((<-chan int)(c))
}

func populateChannel(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}
func pullFromChan(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
