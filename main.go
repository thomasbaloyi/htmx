package main

import (
	"html/template"
	"log"
	"net/http"
)

func root(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "hx-request, hx-target, hx-current-url")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	http.HandleFunc("/add", root(add))
	http.HandleFunc("/edit", root(add))
	http.HandleFunc("/delete", root(add))

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func add(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/add.html")
	if err != nil {
		http.Error(w, "Something went wrong while fetching error template", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	t.Execute(w, nil)
}

func edit(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	t.Execute(w, nil)
}

func delete(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/remove.html")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	t.Execute(w, nil)
}
