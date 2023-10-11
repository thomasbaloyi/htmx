package database

import (
	"log"
	"os"
	"strconv"

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
	n, err = file.WriteString(strconv.Itoa(task.Id) + " " + task.Description + " " + task.CreatedAt + "\n")

	if err != nil {
		return err
	}

	log.Println("Successfully wrote to database file" + strconv.Itoa(n))

	defer file.Close()

	return nil
}
