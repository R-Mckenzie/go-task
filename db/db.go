package db

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Title string    `json:"title"`
	Desc  string    `json:"desc"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

const databasePath = "/.cache/gotask"
const bucketName = "tasks"

func openDb() *bolt.DB {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dbLoc := home + databasePath
	os.MkdirAll(dbLoc, os.ModePerm)
	db, err := bolt.Open(dbLoc+"/tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func LoadTasks() ([]Task, error) {
	db := openDb()
	tasks := []Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.ForEach(func(k, v []byte) error {
			task := decodeJSON(v)
			tasks = append(tasks, task)
			return nil
		})
		return err
	})
	return tasks, err
}

func Save(t Task) error {
	db := openDb()
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		taskID, _ := bucket.NextSequence()
		encoded, err := json.Marshal(&t)
		return bucket.Put([]byte(itob(int(taskID))), encoded)
	})
}

func decodeJSON(jsonStr []byte) Task {
	task := Task{}
	if err := json.Unmarshal(jsonStr, &task); err != nil {
		panic(err)
	}
	return task
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
