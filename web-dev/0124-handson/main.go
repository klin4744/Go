// Create a data structure to pass to a template which
// contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Menu struct {
	Type   string
	Dishes []string
}
type Restaurant struct {
	Name      string
	Breakfast Menu
	Lunch     Menu
	Dinner    Menu
}

func main() {
	restaurants := []Restaurant{
		Restaurant{
			"Chicken Lovers",
			Menu{
				"Breakfast",
				[]string{"1/4 BBQ Ribs", "1/4 Chicken"},
			},
			Menu{
				"Lunch",
				[]string{"1/2 BBQ Ribs", "1/2 Chicken"},
			},
			Menu{
				"Dinner",
				[]string{"Whole BBQ Ribs", "Whole Chicken"},
			},
		},
		Restaurant{
			"McDonalds",
			Menu{
				"Breakfast",
				[]string{"McMuffin", "Big Breakfast"},
			},
			Menu{
				"Lunch",
				[]string{"Chicken Nuggets", "Fish Fillet"},
			},
			Menu{
				"Dinner",
				[]string{"40pc Chicken Nuggets", "Kid's Meal"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
