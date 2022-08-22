package task

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/variables"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup("http://127.0.0.1:8067/flowable-rest/", "rest-admin", "123456", 10086)
	m.Run()
}

//go:embed js3.png
var image1 string

//go:embed js4.png
var image2 string

const id = "e40a9b9b-21b9-11ed-b5b7-0242ac140002"

func TestTasks_List(t *testing.T) {
	var task Tasks
	data, err := task.List(ListRequest{})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_Detail(t *testing.T) {
	var task Tasks
	data, err := task.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_Update(t *testing.T) {
	var task Tasks

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
		data, err := task.Update(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}
}

func TestTasks_Action(t *testing.T) {
	var task Tasks

	type args struct {
		deploymentId string
		req          ActionRequest
	}

	tests := []args{
		//{
		//	deploymentId: id,
		//	req: ActionRequest{
		//		Action: "complete",
		//		Variables: []variables.VariableRequest{
		//			{
		//				Name:  "approved",
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
				Action:   "delegate",
				Assignee: "ceo",
			},
		},
	}
	for _, v := range tests {
		err := task.Action(v.deploymentId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("success")
	}
}

func TestTasks_Delete(t *testing.T) {
	var task Tasks
	err := task.Delete(id, DeleteRequest{
		CascadeHistory: false,
		DeleteReason:   "I like, so what",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}

func TestTasks_ListVariables(t *testing.T) {
	var task Tasks
	data, err := task.ListVariables(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DetailVariables(t *testing.T) {
	var task Tasks
	data, err := task.DetailVariables(id, "myTaskVariable")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_AddVariables(t *testing.T) {
	var task Tasks
	data, err := task.AddVariables(id, []variables.VariableRequest{
		{
			Name:  "myTaskVariable",
			Type:  "string",
			Value: "Hello my friend",
		},
		{
			Name:  "myTaskVariable2",
			Type:  "string",
			Value: "Hello my friend2",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_UpdateVariable(t *testing.T) {
	var task Tasks
	data, err := task.UpdateVariable(id, variables.VariableRequest{
		Name:  "myTaskVariable2",
		Type:  "string",
		Value: "Hello my friend3",
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DeleteVariables(t *testing.T) {
	var task Tasks
	err := task.DeleteVariables(id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}

func TestTasks_DeleteVariable(t *testing.T) {
	var task Tasks
	err := task.DeleteVariable(id, "myTaskVariable")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}

func TestTasks_AddBinaryVariable(t *testing.T) {
	var task Tasks
	data, err := task.AddBinaryVariable(id, variables.FileVariableRequest{
		Name:     "file",
		FileName: "xfxfads",
		Scope:    "",
		Value:    strings.NewReader(image1),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_UpdateBinaryVariable(t *testing.T) {
	var task Tasks
	data, err := task.UpdateBinaryVariable(id, variables.FileVariableRequest{
		Name:     "file",
		FileName: "xfxfads.xxx",
		Scope:    "",
		Value:    strings.NewReader(image2),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DetailBinaryVariable(t *testing.T) {
	var task Tasks
	data, err := task.DetailBinaryVariable(id, "file")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}

func TestTasks_ListIdentity(t *testing.T) {
	var task Tasks
	data, err := task.ListIdentity(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_ListUsersIdentity(t *testing.T) {
	var task Tasks
	data, err := task.ListUsersIdentity(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_ListGroupsIdentity(t *testing.T) {
	var task Tasks
	data, err := task.ListGroupsIdentity(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DetailIdentity(t *testing.T) {
	var task Tasks

	data, err := task.DetailIdentity(id, DetailIdentityRequest{
		Family:     "users",
		IdentityId: "Bobo",
		Type:       "assignee",
	})

	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_AddIdentity(t *testing.T) {
	var task Tasks

	data, err := task.AddIdentity(id, AddIdentityRequest{
		User:  "bobo2",
		Group: "",
		Type:  "myType",
	})

	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DeleteIdentity(t *testing.T) {
	var task Tasks

	err := task.DeleteIdentity(id, DetailIdentityRequest{
		Family:     "users",
		IdentityId: "bobo2",
		Type:       "myType",
	})

	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}
