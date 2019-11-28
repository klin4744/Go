package main

import (
	"fmt"
)

type set int

var x set

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
