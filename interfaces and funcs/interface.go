package main

import (
	"fmt"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}
func (c Circle) getArea() float64 {
	return c.radius * c.radius * 3.14
}

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	square := Square{
		side: 4,
	}
	printArea(square)
	circle := Circle{
		radius: 4,
	}
	printArea(circle)

}
