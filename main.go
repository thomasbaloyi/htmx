package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var COUNT = 1

type Task struct {
	Id          int
	Description string
}

var tasks = []Task{}

func main() {

	http.HandleFunc("/add", corsHandler(add))

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func corsHandler(next http.HandlerFunc) http.HandlerFunc {
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

	tasks = append(tasks, Task{Id: COUNT, Description: strings.Split(string(body), "=")[1]})
	COUNT++

	t := ""

	for _, task := range tasks {
		t = t + "<li>" + strconv.Itoa(task.Id) + " - " + task.Description + "</li>"
	}

	tmp, err := template.New("tasks").Parse(t)
	if err != nil {
		log.Panic(err)
	}

	tmp.Execute(w, nil)

}
