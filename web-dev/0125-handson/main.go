package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Stock struct {
	Date string
	High string
	Low  string
}

func main() {
	csvfile, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	r := csv.NewReader(csvfile)
	data, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data[1])
	stocks := []Stock{}
	for row := 1; row < len(data); row++ {
		currentStock := Stock{data[row][0], data[row][2], data[row][3]}
		stocks = append(stocks, currentStock)
	}
	err = tpl.Execute(os.Stdout, stocks)
	if err != nil {
		log.Fatalln(err)
	}
}
