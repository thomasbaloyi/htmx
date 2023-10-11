package database

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"example.com/htmx/model"
)

var filename string = "./database/db.txt"

func CreateDbFile() bool {

	file, err := os.OpenFile(filename, os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer file.Close()

	log.Printf("Successfully created database file")

	return true
}

func WriteToDbFile(task model.Task) error {
	file, err := os.OpenFile(filename, os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	n := 0
	n, err = file.WriteString(strconv.Itoa(task.Id) + ", " + task.Description + ", " + task.CreatedAt + "\n")

	if err != nil {
		return err
	}

	log.Println("Successfully wrote to database file" + strconv.Itoa(n))

	defer file.Close()

	return nil
}

func ReadFromDbFile() ([]model.Task, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	tasks := []model.Task{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		taskList := strings.Split(scanner.Text(), ",")
		id, err := strconv.Atoi(taskList[0])
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, model.Task{Description: taskList[1], Id: id, CreatedAt: taskList[2]})
	}

	return tasks, nil
}
