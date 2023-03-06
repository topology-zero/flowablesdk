package form_deployment

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/external_form/model"
)

const id = "54a0dca1-3f08-11ed-a7b8-38f3ab6b92c1"

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

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

func TestAdd(t *testing.T) {
	data, err := Add(AddRequest{
		FileName: "test1.form",
		Data: model.Model{
			Name: "请假2流程",
			Key:  "form2",
			Fields: []model.FormField{
				{
					Id:   "startTime",
					Name: "开始时间1",
				},
				{
					Id:   "days",
					Name: "请假天数1",
				},
				{
					Id:   "reason",
					Name: "请假原因1",
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
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

func TestDelete(t *testing.T) {
	err := Delete(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}
