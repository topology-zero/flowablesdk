package process_instance

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/candidate"
	"github.com/MasterJoyHunan/flowablesdk/variable"
)

//go:embed js.png
var png string

//go:embed js2.png
var png2 string

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const id = "1786d504-b255-11ed-b3e2-38f3ab6b92c1"

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
	data, err := Update(id, UpdateRequest{
		Action: "suspend",
	})

	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))

	data, err = Update(id, UpdateRequest{
		Action: "activate",
	})

	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ = json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	err := Delete(id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete sucess")
}

func TestStart(t *testing.T) {
	data, err := Start(StartProcessRequest{
		ProcessDefinitionId:  "",
		ProcessDefinitionKey: "holidayRequest",
		Message:              "",
		BusinessKey:          "myBusinessKey",
		Variables:            nil,
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestListCandidate(t *testing.T) {
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
		InstanceId string
		req        candidate.Candidate
	}
	tests := []args{
		{
			InstanceId: id,
			req: candidate.Candidate{
				User: "joy",
				Type: "CEO",
			},
		},
		{
			InstanceId: id,
			req: candidate.Candidate{
				User: "BOBO",
				Type: "CEO",
			},
		},
	}
	for _, v := range tests {
		data, err := AddCandidate(v.InstanceId, v.req)
		if err != nil {
			t.Error(err)
			return
		}

		jsonStr, _ := json.MarshalIndent(&data, "", "    ")
		fmt.Println(string(jsonStr))
	}
}

func TestDeleteCandidate(t *testing.T) {
	type args struct {
		InstanceId string
		req        candidate.Candidate
	}
	tests := []args{
		{
			InstanceId: id,
			req: candidate.Candidate{
				User: "BOBO",
				Type: "CEO",
			},
		},
		{
			InstanceId: id,
			req: candidate.Candidate{
				User: "joy",
				Type: "CEO",
			},
		},
	}
	for _, v := range tests {
		err := DeleteCandidate(v.InstanceId, v.req)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("delete success")
	}
}

func TestListVariables(t *testing.T) {
	data, err := ListVariables(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAddVariables(t *testing.T) {
	data, err := AddVariables(id, []variable.VariableRequest{
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

func TestUpdateVariables(t *testing.T) {
	data, err := UpdateVariables(id, []variable.VariableRequest{
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

func TestUpdateVariable(t *testing.T) {
	data, err := UpdateVariable(id, "XXX", variable.VariableRequest{
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

func TestVariableDetail(t *testing.T) {
	data, err := VariableDetail(id, "XXX")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAddFileVariable(t *testing.T) {
	data, err := AddFileVariable(id, variable.FileVariableRequest{
		VariableName: "costomFile2",
		FileName:     "js.png",
		Value:        strings.NewReader(png),
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestUpdateFileVariable(t *testing.T) {
	data, err := UpdateFileVariable(id, variable.FileVariableRequest{
		VariableName: "costomFile",
		FileName:     "js.png",
		Value:        strings.NewReader(png2),
	})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
