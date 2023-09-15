package main

import (
	"html/template"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	log.Println("user accessed")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}

func main() {

	http.HandleFunc("/", root)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
