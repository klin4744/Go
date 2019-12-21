package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))
	err := tpl.ExecuteTemplate(w, "home.gohtml", "Kevin")
	if err != nil {
		log.Fatalln(err)
	}
}
func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Dog Route")
}

func name(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Kevin Lin")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/name", name)
	http.ListenAndServe(":8080", nil)
}
