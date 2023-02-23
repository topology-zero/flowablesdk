package history_process_instance

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const id = "58ad0c6e-b25a-11ed-b3e2-38f3ab6b92c1"

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
	req := ListRequest{}
	req.Start = 0
	data, err := Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	req := ListRequest{}
	req.Start = 0
	err := Delete(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestListIdentity(t *testing.T) {
	req := ListRequest{}
	req.Start = 0
	data, err := ListIdentity(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestBinaryData(t *testing.T) {
	req := ListRequest{}
	req.Start = 0
	data, err := BinaryData(id, "")
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
