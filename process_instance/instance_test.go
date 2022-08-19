package process_instance

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

//go:embed js.png
var png string

//go:embed js2.png
var png2 string

func TestMain(m *testing.M) {
	flowablesdk.Setup("http://127.0.0.1:8067/flowable-rest/", "rest-admin", "123456", 10086)
	m.Run()
}

const id = "9d5a8b48-1f68-11ed-a343-0242ac140002"

func TestInstance_List(t *testing.T) {
	var d Instance
	data, err := d.List(ListRequest{})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_Detail(t *testing.T) {
	var d Instance
	data, err := d.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_Update(t *testing.T) {
	var d Instance
	data, err := d.Update(id, UpdateRequest{
		Action: "suspend",
	})

	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))

	data, err = d.Update(id, UpdateRequest{
		Action: "activate",
	})

	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ = json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_Delete(t *testing.T) {
	var d Instance
	err := d.Delete(id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete sucess")
}

func TestInstance_Start(t *testing.T) {
	var d Instance
	data, err := d.Start(StartProcessRequest{
		ProcessDefinitionId:  "",
		ProcessDefinitionKey: "holidayRequest",
		Message:              "",
		BusinessKey:          "myBusinessKey",
		ReturnVariables:      false,
		Variables:            nil,
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_ListCandidate(t *testing.T) {
	var d Instance
	data, err := d.ListCandidate(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_AddCandidate(t *testing.T) {
	var d Instance
	type args struct {
		InstanceId string
		req        Candidate
	}
	tests := []args{
		{
			InstanceId: id,
			req: Candidate{
				User: "joy",
				Type: "CEO",
			},
		},
		{
			InstanceId: id,
			req: Candidate{
				User: "BOBO",
				Type: "CEO",
			},
		},
	}
	for _, v := range tests {
		data, err := d.AddCandidate(v.InstanceId, v.req)
		if err != nil {
			t.Error(err)
			return
		}

		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}
}

func TestInstance_DeleteCandidate(t *testing.T) {
	var d Instance
	type args struct {
		InstanceId string
		req        Candidate
	}
	tests := []args{
		{
			InstanceId: id,
			req: Candidate{
				User: "BOBO",
				Type: "CEO",
			},
		},
		{
			InstanceId: id,
			req: Candidate{
				User: "joy",
				Type: "CEO",
			},
		},
	}
	for _, v := range tests {
		err := d.DeleteCandidate(v.InstanceId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("delete success")
	}
}

func TestInstance_ListVariables(t *testing.T) {
	var d Instance
	data, err := d.ListVariables(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_AddVariables(t *testing.T) {
	var d Instance
	data, err := d.AddVariables(id, []VariableRequest{
		{
			Name:  "XXX",
			Type:  "string",
			Value: "1123",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_EditVariables(t *testing.T) {
	var d Instance
	data, err := d.EditVariables(id, []VariableRequest{
		{
			Name:  "XXX",
			Type:  "integer",
			Value: 1123,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_EditVariable(t *testing.T) {
	var d Instance
	data, err := d.EditVariable(id, "XXX", VariableRequest{
		Name:  "XXX",
		Type:  "integer",
		Value: 11238,
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_VariableDetail(t *testing.T) {
	var d Instance
	data, err := d.VariableDetail(id, "XXX")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_AddFileVariable(t *testing.T) {
	var d Instance
	data, err := d.AddFileVariable(id, FileVariableRequest{
		Name:     "costomFile2",
		FileName: "js.png",
		Value:    strings.NewReader(png),
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestInstance_EditFileVariable(t *testing.T) {
	var d Instance
	data, err := d.EditFileVariable(id, FileVariableRequest{
		Name:     "costomFile",
		FileName: "js.png",
		Value:    strings.NewReader(png2),
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
