package task_event

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/topology-zero/flowablesdk"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const (
	id       = "95b36aeb-b420-11ed-b7d1-38f3ab6b92c1"
	eventsId = "b57f1602-b34a-11ed-b7d1-38f3ab6b92c1"
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
	data, err := Detail(id, eventsId)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	err := Delete(id, eventsId)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}
