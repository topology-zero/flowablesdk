package task

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const id = "17891efa-b255-11ed-b3e2-38f3ab6b92c1"

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
				Assignee:        "Bobo",
				DelegationState: "",
				Description:     "",
				DueDate:         "",
				Name:            "",
				Owner:           "",
				ParentTaskId:    "",
				Priority:        0,
			},
		},
		{
			deploymentId: id,
			req: UpdateRequest{
				Assignee:        "",
				DelegationState: "",
				Description:     "xxxxx",
				DueDate:         "",
				Name:            "",
				Owner:           "",
				ParentTaskId:    "",
				Priority:        0,
			},
		},
		{
			deploymentId: id,
			req: UpdateRequest{
				Assignee:        "",
				DelegationState: "",
				Description:     "",
				DueDate:         "",
				Name:            "",
				Owner:           "master bobo",
				ParentTaskId:    "",
				Priority:        0,
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

func TestAction(t *testing.T) {
	type args struct {
		deploymentId string
		req          ActionRequest
	}

	tests := []args{
		//{
		//	deploymentId: id,
		//	req: ActionRequest{
		//		Action: "complete",
		//		Variables: []variable.VariableRequest{
		//			{
		//				VariableName:  "approved",
		//				Type:  "boolean",
		//				Value: true,
		//			},
		//		},
		//	},
		//},
		//{
		//	deploymentId: id,
		//	req: ActionRequest{
		//		Action:   "claim",
		//		Assignee: "userWhoClaims",
		//	},
		//},
		{
			deploymentId: id,
			req: ActionRequest{
				Action: "complete",
				//Assignee: "ceo",
			},
		},
	}
	for _, v := range tests {
		err := Action(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("success")
	}
}

func TestDelete(t *testing.T) {
	err := Delete(id, DeleteRequest{
		CascadeHistory: false,
		DeleteReason:   "I like, so what",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}
