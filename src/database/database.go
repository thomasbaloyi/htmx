package database

import (
	"log"
	"os"
)

func CreateDbFile() bool {
	name := "./database/db.txt"

	file, err := os.OpenFile(name, os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
		return false
	}

	file.Close()

	log.Printf("Successfully created database file")

	return true
}
