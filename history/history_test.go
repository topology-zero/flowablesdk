package history

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

const id = "3cb5892b-3571-11ed-a314-38f3ab6b92c1"

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{
		Url: "http://127.0.0.1:8080/",
		GetUserId: func() string {
			return "1"
		},
	})
	m.Run()
}

func TestHistory_List(t *testing.T) {
	var his History
	req := ListRequest{}
	data, err := his.List(req)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestHistory_Detail(t *testing.T) {
	var his History
	req := ListRequest{}
	req.Start = 0
	data, err := his.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestHistory_Delete(t *testing.T) {
	var his History
	req := ListRequest{}
	req.Start = 0
	err := his.Delete(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestHistory_ListIdentity(t *testing.T) {
	var his History
	req := ListRequest{}
	req.Start = 0
	data, err := his.ListIdentity(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
