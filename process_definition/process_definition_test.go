package process_definition

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

const id = "Process_1677032402556:1:7c716b23-b257-11ed-b3e2-38f3ab6b92c1"

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

func TestUpdate(t *testing.T) {

	type args struct {
		deploymentId string
		req          UpdateRequest
	}

	tests := []args{
		{
			deploymentId: id,
			req: UpdateRequest{
				Category: "x1",
			},
		},
		{
			deploymentId: id,
			req: UpdateRequest{
				Action: "suspend",
			},
		},
		{
			deploymentId: id,
			req: UpdateRequest{
				Action: "activate",
			},
		},
	}
	for _, v := range tests {
		data, err := Update(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}

}

func TestResourceContent(t *testing.T) {
	data, err := ResourceContent(id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}

func TestModel(t *testing.T) {
	data, err := Model(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestCandidate(t *testing.T) {
	data, err := ListCandidate(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAddCandidate(t *testing.T) {

	type args struct {
		deploymentId string
		req          AddCandidateRequest
	}

	tests := []args{
		{
			deploymentId: id,
			req: AddCandidateRequest{
				User:  "",
				Group: "bobo",
			},
		},
		{
			deploymentId: id,
			req: AddCandidateRequest{
				User:  "joy",
				Group: "",
			},
		},
	}
	for _, v := range tests {
		data, err := AddCandidate(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}
}

func TestDeleteCandidate(t *testing.T) {
	err := DeleteCandidate(id, DeleteCandidateRequest{
		Family:      "users",
		CandidateId: "ajax",
	})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestCandidateDetail(t *testing.T) {
	data, err := CandidateDetail(id, DeleteCandidateRequest{
		Family:      "users",
		CandidateId: "ajax",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
