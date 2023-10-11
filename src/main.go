package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"example.com/htmx/database"

	"example.com/htmx/model"
)

var COUNT = 1

var tasks = []model.Task{}

func main() {

	isDbFileCreated := database.CreateDbFile()

	if isDbFileCreated {
		log.Printf("DB file created successfully, application starting...")
	} else {
		log.Panic("Failed to create database file")
	}

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	fsImages := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", fsImages))

	http.HandleFunc("/", corsHandler(rootHandler))
	http.HandleFunc("/add", corsHandler(add))
	http.HandleFunc("/tasks", corsHandler(taskHandler))
	http.HandleFunc("/home", corsHandler(homeHandler))
	http.HandleFunc("/task/add", corsHandler(addHandler))
	http.HandleFunc("/update", corsHandler(editHandler))

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

func taskHandler(w http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("templates/list.html")
	if err != nil {
		log.Panic(err)
	}

	tmp.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Panic(err)
	}

	tmp.Execute(w, nil)
}

func add(w http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("templates/add.html")
	if err != nil {
		log.Panic(err)
	}

	tmp.Execute(w, nil)
}

func addHandler(w http.ResponseWriter, req *http.Request) {
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
	} else {
		description, err = url.QueryUnescape(description)
		if err != nil {
			log.Panic(err)
		}
	}

	tasks = append(tasks, model.Task{Id: COUNT, Description: description, CreatedAt: time.Now().Format("2006-01-02 15:04")})
	COUNT++

	tmp, err := template.ParseFiles("templates/list.html")
	if err != nil {
		log.Panic(err)
	}

	tmp.Execute(w, tasks)

}

func editHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("Id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id = id - 1

	task := tasks[id]

	tmp, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		log.Panic(err)
	}

	tmp.Execute(w, task)
}
