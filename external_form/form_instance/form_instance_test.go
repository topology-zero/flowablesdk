package form_instance

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const (
	id     = "c75fd0e2-b359-11ed-b7d1-38f3ab6b92c1"
	taskId = "73243e27-b256-11ed-b3e2-38f3ab6b92c1"
	key    = "form3"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

func TestList(t *testing.T) {
	data, count, err := List(ListRequest{})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
	fmt.Println()
	fmt.Println(count)
}

func TestAdd(t *testing.T) {
	req := AddRequest{
		FormDefinitionKey: key,
		TaskId:            taskId,
		Variables: map[string]any{
			"days": "166",
		},
	}
	err := Add(req)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("add success")
}

func TestDetail(t *testing.T) {
	data, err := Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestModel(t *testing.T) {
	data, err := Model(AddRequest{
		FormInstanceId:    id,
		FormDefinitionKey: key,
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDetailAndModel(t *testing.T) {
	data, err := DetailAndModel(AddRequest{
		FormInstanceId:    id,
		FormDefinitionKey: key,
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
