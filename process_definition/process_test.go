package process_definition

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup("http://127.0.0.1:8067/flowable-rest/", "rest-admin", "123456", 10086)
	m.Run()
}

const id = "holidayRequest:1:95a702e4-21b9-11ed-b5b7-0242ac140002"

func TestProcess_List(t *testing.T) {
	var d Process
	data, err := d.List(ListRequest{})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestProcess_Detail(t *testing.T) {
	var d Process
	data, err := d.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestProcess_Update(t *testing.T) {
	var d Process

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
		data, err := d.Update(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}

}

func TestProcess_ResourceContent(t *testing.T) {
	var d Process
	data, err := d.ResourceContent(id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}

func TestProcess_Model(t *testing.T) {
	var d Process
	data, err := d.Model(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestProcess_Candidate(t *testing.T) {
	var d Process
	data, err := d.ListCandidate(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestProcess_AddCandidate(t *testing.T) {
	var d Process

	type args struct {
		deploymentId string
		req          AddCandidateRequest
	}

	tests := []args{
		//{
		//	deploymentId: id,
		//	req: AddCandidateRequest{
		//		User:    "",
		//		Group: "bobo",
		//	},
		//},
		{
			deploymentId: id,
			req: AddCandidateRequest{
				User:    "joy",
				GroupId: "",
			},
		},
		{
			deploymentId: id,
			req: AddCandidateRequest{
				User:    "ajax",
				GroupId: "ceo",
			},
		},
	}
	for _, v := range tests {
		data, err := d.AddCandidate(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}
}

func TestProcess_DeleteCandidate(t *testing.T) {
	var d Process
	err := d.DeleteCandidate(id, DeleteCandidateRequest{
		Family:      "users",
		CandidateId: "ajax",
	})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestProcess_CandidateDetail(t *testing.T) {
	var d Process
	data, err := d.CandidateDetail(id, DeleteCandidateRequest{
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
