package deployment

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const id = "54a0dca1-3f08-11ed-a7b8-38f3ab6b92c1"

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

func TestExternalFormDeployment_List(t *testing.T) {
	f := new(ExternalFormDeployment)
	req := ListRequest{}
	data, err := f.List(req)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestExternalFormDeployment_Add(t *testing.T) {
	f := new(ExternalFormDeployment)
	data, err := f.Add(AddRequest{
		FileName: "test.form",
		Category: "xxxxx",
		Data: `{
"key": "form3",
"name": "请假2流程",
"fields": [
            {
            "id": "startTime",
            "name": "开始时间1",
            "type": "date",
            "required": true,
            "placeholder": "empty"
            },
            {
            "id": "days",
            "name": "请假天数1",
            "type": "string",
            "required": true,
            "placeholder": "empty"
            },
            {
            "id": "reason",
            "name": "请假原因1",
            "type": "text",
            "required": true,
            "placeholder": "empty"
            }
    ]
}`,
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestExternalFormDeployment_Detail(t *testing.T) {
	f := new(ExternalFormDeployment)
	data, err := f.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestExternalFormDeployment_Delete(t *testing.T) {
	f := new(ExternalFormDeployment)
	err := f.Delete(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}
