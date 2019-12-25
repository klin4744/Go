package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Dog struct {
	Name  string
	Breed string
	Age   int
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("dog.gohtml"))
	myDog := Dog{"Simon", "Husky", 2}
	err := tpl.Execute(w, myDog)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
