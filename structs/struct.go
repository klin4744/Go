package main

import (
	"fmt"
)

type person struct {
	firstName       string
	lastName        string
	favoriteFlavors []string
}
type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := person{
		firstName:       "Kevin",
		lastName:        "Lin",
		favoriteFlavors: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		firstName:       "Bob",
		lastName:        "Tom",
		favoriteFlavors: []string{"cookies and cream", "vanilla"},
	}
	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.favoriteFlavors {
		fmt.Printf("\t%v\n", v)
	}
	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.favoriteFlavors {
		fmt.Printf("\t%v\n", v)
	}
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favoriteFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-----------------")
	}
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(t)
	fmt.Println(s)

}
