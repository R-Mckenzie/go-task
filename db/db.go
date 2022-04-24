package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Id    uint64
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

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
	return db
}

func LoadTasks() ([]Task, error) {
	db := openDb()
	defer db.Close()
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
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		taskID, _ := bucket.NextSequence()
		t.Id = taskID
		fmt.Printf("%v\n", taskID)
		encoded, err := json.Marshal(&t)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(itob(t.Id)), encoded)
	})
}

func Delete(id int) error {
	db := openDb()
	defer db.Close()
	fmt.Printf("%v\n", id)
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		count := 0
		return b.ForEach(func(k, v []byte) error {
			if count == id-1 {
				b.Delete(k)
				return nil
			}
			count++
            return fmt.Errorf("Task to delete not found")
		})
	})
}

func decodeJSON(jsonStr []byte) Task {
	task := Task{}
	if err := json.Unmarshal(jsonStr, &task); err != nil {
		panic(err)
	}
	return task
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
