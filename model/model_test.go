package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const id = "b3aa9242-b410-11ed-b7d1-38f3ab6b92c1"

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

func TestDetail(t *testing.T) {
	data, err := Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAdd(t *testing.T) {
	data, err := Add(AddRequest{
		Name: "TestModel",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestUpdate(t *testing.T) {
	data, err := Update(id, AddRequest{
		Name: "TestModel3",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	err := Delete(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}