package definition

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const id = "54b21ab3-3f08-11ed-a7b8-38f3ab6b92c1"

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

func TestExternalFormDefinition_List(t *testing.T) {
	f := new(ExternalFormDefinition)
	req := ListRequest{}
	data, err := f.List(req)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestExternalFormDefinition_Detail(t *testing.T) {
	f := new(ExternalFormDefinition)
	data, err := f.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestExternalFormDeployment_Model(t *testing.T) {
	f := new(ExternalFormDefinition)
	data, err := f.Model(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
