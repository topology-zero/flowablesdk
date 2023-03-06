package task_comment

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/comment"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const (
	id        = "73243e27-b256-11ed-b3e2-38f3ab6b92c1"
	commentId = "120ae147-b34d-11ed-b7d1-38f3ab6b92c1"
)

func TestList(t *testing.T) {
	data, err := List(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDetail(t *testing.T) {
	data, err := Detail(id, commentId)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAdd(t *testing.T) {
	data, err := Add(id, comment.AddComment{
		Message:               "user comments v2",
		SaveProcessInstanceId: false,
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	err := Delete(id, commentId)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}
