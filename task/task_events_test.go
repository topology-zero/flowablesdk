package task

import (
	"encoding/json"
	"fmt"
	"testing"
)

const eventsId = "f511dc15-21ec-11ed-b5b7-0242ac140002"

func TestTasks_ListEvents(t *testing.T) {
	var task Tasks
	data, err := task.ListEvents(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DetailEvents(t *testing.T) {
	var task Tasks
	data, err := task.DetailEvents(id, eventsId)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
