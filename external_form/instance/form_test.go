package instance

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const (
	id  = "ed9c595f-3fcf-11ed-9d90-38f3ab6b92c1"
	key = "form2"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

func TestExternalFormInstance_List(t *testing.T) {
	f := new(ExternalFormInstance)
	req := ListRequest{}
	data, err := f.List(req)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestExternalFormInstance_Add(t *testing.T) {
	f := new(ExternalFormInstance)
	req := AddRequest{
		FormDefinitionKey:   key,
		TaskId:              "d46d74ea-3ed6-11ed-9508-38f3ab6b92c1",
		ProcessDefinitionId: "",
		Variables: map[string]any{
			"days": "166",
		},
	}
	err := f.Add(req)
	if err != nil {
		t.Error(err)
		return
	}

	//req = AddRequest{
	//	FormDefinitionId:    "",
	//	TaskId:              "",
	//	ProcessDefinitionId: "",
	//	Variables: map[string]any{
	//		"": "",
	//	},
	//}
	//err = f.Add(req)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}

	fmt.Println("add success")
}

func TestExternalFormInstance_Detail(t *testing.T) {
	f := new(ExternalFormInstance)
	data, err := f.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestExternalFormInstance_Model(t *testing.T) {
	f := new(ExternalFormInstance)
	data, err := f.Model(AddRequest{
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

func TestExternalFormInstance_DetailAndModel(t *testing.T) {
	f := new(ExternalFormInstance)
	data, err := f.DetailAndModel(AddRequest{
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
