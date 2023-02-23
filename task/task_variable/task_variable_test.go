package task_variable

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

//go:embed _js3.png
var image1 []byte

//go:embed _js4.png
var image2 []byte

const id = "73243e27-b256-11ed-b3e2-38f3ab6b92c1"

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
	data, err := Detail(id, "process_instance_var")
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDetailBinary(t *testing.T) {
	data, err := DetailBinary(id, "file")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(data))
}

func TestAdd(t *testing.T) {
	data, err := Add(id, []variables.VariableRequest{
		{
			Name:  "test1",
			Type:  "string",
			Value: "hello ",
		},
		{
			Name:  "test2",
			Type:  "string",
			Value: "world",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAddBinary(t *testing.T) {
	data, err := AddBinary(id, variables.FileVariableRequest{
		VariableName: "file",
		FileName:     "_js3.png",
		Value:        bytes.NewReader(image1),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestUpdate(t *testing.T) {
	data, err := Update(id, variables.VariableRequest{
		Name:  "test1",
		Type:  "string",
		Value: "fuck ",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestUpdateBinary(t *testing.T) {
	data, err := UpdateBinary(id, variables.FileVariableRequest{
		VariableName: "file",
		FileName:     "_js4.png",
		Value:        bytes.NewReader(image2),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDeleteAllLocal(t *testing.T) {
	err := DeleteAllLocal(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestDelete(t *testing.T) {
	err := Delete(id, "test1")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}
