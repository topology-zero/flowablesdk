package task_form

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

const id = "728fd533-b9c8-11ed-ac5d-38f3ab6b92c1"

func TestGetForm(t *testing.T) {
	data, err := GetForm(id)

	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
