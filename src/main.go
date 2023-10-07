package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var COUNT = 1

type Task struct {
	Id          int
	Description string
	CreatedAt   time.Time
}

var tasks = []Task{}

func main() {

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", corsHandler(rootHandler))
	http.HandleFunc("/add", corsHandler(add))

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")

	if err != nil {
		log.Panic(err)
		http.Error(w, "Something went wrong generating response.", http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}

func corsHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method, r.URL, r.RemoteAddr)

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

func add(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(w, "Could not extract payload from request", http.StatusInternalServerError)
		log.Panic(err)
	}

	defer req.Body.Close()

	description := strings.Split(string(body), "=")[1]

	if len(description) == 0 {
		w.Header().Set("Access-Control-Allow-Headers", "HX-Retarget")
		http.Error(w, "No description.", http.StatusPreconditionFailed)
		return
	}

	tasks = append(tasks, Task{Id: COUNT, Description: description, CreatedAt: time.Now()})
	COUNT++

	tmp, err := template.ParseFiles("templates/list.html")
	if err != nil {
		log.Panic(err)
	}

	tmp.Execute(w, tasks)

}
