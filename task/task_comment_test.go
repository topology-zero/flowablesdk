package task

import (
	"encoding/json"
	"fmt"
	"testing"
)

const commentId = "f7477967-21ef-11ed-b5b7-0242ac140002"

func TestTasks_ListComments(t *testing.T) {
	var task Tasks
	data, err := task.ListComments(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DetailComments(t *testing.T) {
	var task Tasks
	data, err := task.DetailComments(id, commentId)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_AddComments(t *testing.T) {
	var task Tasks
	data, err := task.AddComments(id, AddCommentsRequest{
		Message:               "user comments",
		SaveProcessInstanceId: true,
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DeleteComments(t *testing.T) {
	var task Tasks
	err := task.DeleteComments(id, commentId)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}
