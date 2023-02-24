package execution

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const id = "d8c27ea9-b256-11ed-b3e2-38f3ab6b92c1"

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

func TestList(t *testing.T) {
	req := ListRequest{}
	data, _, err := List(req)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
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

func TestListActive(t *testing.T) {
	active, err := ListActive(id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(active)
}

func TestExecute(t *testing.T) {
	active, err := Execute(id, ExecuteRequest{
		Action:      "signal",
		MessageName: "",
		Variables:   nil,
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(active)
}
