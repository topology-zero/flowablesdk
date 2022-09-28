package form

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const id = "Process_1664178530144:1:17b0612b-3d70-11ed-a6e0-38f3ab6b92c1"

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

func TestForm_GetFrom(t *testing.T) {
	f := new(Form)
	data, err := f.GetFrom(GetFromRequest{
		ProcessDefinitionId: id,
	})

	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestForm_SubmitForm(t *testing.T) {
	f := new(Form)
	data, err := f.SubmitForm(SubmitFromRequest{
		ProcessDefinitionId: id,
		Properties: []Properties{
			{
				Id:    "xxx",
				Value: "xyz",
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
