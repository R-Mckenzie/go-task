package db

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	// INTEGRATION TESTS
	task := Task{Title: "Test task", Desc: "Description"}
	Save(task)
	tasks, err := LoadTasks()

	if err != nil {
		t.Fatal("did not expect error")
	}

	if tasks[len(tasks)-1] != task {
		t.Errorf("expected %v, got %v", task, tasks[0])
	}
}

func TestDecode(t *testing.T) {
	task := Task{Title: "Test task", Desc: "Description"}
	jsonStr, _ := json.Marshal(task)
	fmt.Printf("json: %s", jsonStr)
	result := decodeJSON(jsonStr)
	fmt.Printf("result: %s", result)

	if result != task {
		t.Errorf("wanted %+v, got %+v", task, result)
	}
}
