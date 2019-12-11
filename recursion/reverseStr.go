package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseString("My name is bob"))

}

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}
