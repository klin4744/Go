package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

// Create a data structure to pass to a template which
// contains information about California hotels including Name, Address, City, Zip, Region
// region can be: Southern, Central, Northern
// can hold an unlimited number of hotels
type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	hotels := []Hotel{
		Hotel{"Super Hotel 1", "Some address in california", "California", 12345, "Southern"},
		Hotel{"Super Hotel 2", "Some address in california 2", "California", 12345, "Northern"},
	}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
