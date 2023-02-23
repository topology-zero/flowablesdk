package task_identity

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

func TestListUsers(t *testing.T) {
	data, err := ListUsers(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestListGroups(t *testing.T) {
	data, err := ListGroups(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDetail(t *testing.T) {
	data, err := Detail(id, DetailRequest{
		Family:     "users",
		IdentityId: "BOBO",
		Type:       "CTO",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAdd(t *testing.T) {
	data, err := Add(id, AddRequest{
		User: "BOBO",
		Type: "CTO",
	})
	if err != nil {
		t.Error(err)
		return
	}

	data, err = Add(id, AddRequest{
		Group: "JOY",
		Type:  "CEO",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	err := Delete(id, DetailRequest{
		Family:     "users",
		IdentityId: "BOBO",
		Type:       "CTO",
	})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}
