package deployment

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

//go:embed test.xml
var xmlStr string

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const id = "291ac182-b323-11ed-b7d1-38f3ab6b92c1"

// 获取列表
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

// 获取详情
func TestDetail(t *testing.T) {
	data, err := Detail(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

// 获取
func TestCreate(t *testing.T) {
	req := CreateRequest{
		FileName: "TestDeployment_Create.bpmn20.xml",
		Xml:      xmlStr,
	}
	data, err := Create(req)
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

func TestResource(t *testing.T) {
	data, err := Resource(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestResourceDetail(t *testing.T) {
	data, err := ResourceDetail(id, "TestDeployment_Create.bpmn20.xml")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestResourceContent(t *testing.T) {
	data, err := ResourceContent(id, "TestDeployment_Create.bpmn20.xml")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}
