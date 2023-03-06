package execution_variable

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/variable"
)

const id = "d8c27ea9-b256-11ed-b3e2-38f3ab6b92c1"

//go:embed _js3.png
var image []byte

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

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

func TestAdd(t *testing.T) {
	data, err := Add(id, []variable.VariableRequest{
		{
			Name:  "a",
			Type:  "integer",
			Value: 1,
		},
		{
			Name:  "b",
			Type:  "integer",
			Value: 2,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestUpdate(t *testing.T) {
	data, err := Update(id, []variable.VariableRequest{
		{
			Name:  "a",
			Type:  "integer",
			Value: 2,
		},
		{
			Name:  "b",
			Type:  "integer",
			Value: 1,
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
	data, err := AddBinary(id, variable.FileVariableRequest{
		VariableName: "image666",
		FileName:     "image1.png",
		Value:        bytes.NewReader(image),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestUpdateBinary(t *testing.T) {
	data, err := UpdateBinary(id, "image666", variable.FileVariableRequest{
		VariableName: "image666",
		FileName:     "image2.png",
		Value:        bytes.NewReader(image),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	err := Delete(id, "image666")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestDeleteAllLocal(t *testing.T) {
	err := DeleteAllLocal(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestBinaryData(t *testing.T) {
	data, err := BinaryData(id, "image666")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(data)
}
