package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h1>Welcome Home</h1>")
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
