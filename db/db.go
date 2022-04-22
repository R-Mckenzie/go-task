package db

import (
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	title string
	desc  string
	start time.Time
	end   time.Time
}

const databasePath = "/.cache/gotask/tasks.db"

func openDb() bolt.DB {
	dbLoc, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dbLoc += databasePath

	db, err := bolt.Open(dbLoc, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func LoadTasks() ([]Task, err) {

}

func SaveTask(task Task) err {

}
