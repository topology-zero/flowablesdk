package process

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
	data, err := d.Detail("createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002")
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
			deploymentId: "createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002",
			req: UpdateRequest{
				Category: "x1",
			},
		},
		{
			deploymentId: "createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002",
			req: UpdateRequest{
				Action: "suspend",
			},
		},
		{
			deploymentId: "createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002",
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
	data, err := d.ResourceContent("createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}

func TestProcess_Model(t *testing.T) {
	var d Process
	data, err := d.Model("createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002")
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestProcess_ProcessCandidate(t *testing.T) {
	var d Process
	data, err := d.ProcessCandidate("createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002")
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestProcess_ProcessAddCandidate(t *testing.T) {
	var d Process

	type args struct {
		deploymentId string
		req          ProcessAddCandidateRequest
	}

	tests := []args{
		//{
		//	deploymentId: "createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002",
		//	req: ProcessAddCandidateRequest{
		//		User:    "",
		//		GroupId: "bobo",
		//	},
		//},
		{
			deploymentId: "createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002",
			req: ProcessAddCandidateRequest{
				User:    "joy",
				GroupId: "",
			},
		},
		{
			deploymentId: "createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002",
			req: ProcessAddCandidateRequest{
				User:    "ajax",
				GroupId: "ceo",
			},
		},
	}
	for _, v := range tests {
		data, err := d.ProcessAddCandidate(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}
}

func TestProcess_ProcessDeleteCandidate(t *testing.T) {
	var d Process
	err := d.ProcessDeleteCandidate("createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002", ProcessDeleteCandidateRequest{
		Family:     "users",
		IdentityId: "ajax",
	})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestProcess_CandidateProcess(t *testing.T) {
	var d Process
	data, err := d.CandidateProcess("createTimersProcess:1:8d3f4a4f-1dfd-11ed-ab65-0242ac170002", ProcessDeleteCandidateRequest{
		Family:     "users",
		IdentityId: "ajax",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
