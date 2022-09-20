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
	flowablesdk.Setup(flowablesdk.Config{
		Url: "http://127.0.0.1:8080/",
		GetUserId: func() string {
			return "1"
		},
	})
	m.Run()
}

const id = "9501976b-3011-11ed-99ce-38f3ab6b92c1"

// 获取列表
func TestDeployment_List(t *testing.T) {
	var d Deployment
	data, err := d.List(ListRequest{})
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

// 获取详情
func TestDeployment_Detail(t *testing.T) {
	var d Deployment
	data, err := d.Detail(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

// 获取
func TestDeployment_Create(t *testing.T) {
	var d Deployment
	req := CreateRequest{
		FileName: "TestDeployment_Create.bpmn20.xml",
		Xml:      xmlStr,
	}
	data, err := d.Create(req)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDeployment_Delete(t *testing.T) {
	var d Deployment
	err := d.Delete(id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}

func TestDeployment_Resource(t *testing.T) {
	var d Deployment
	data, err := d.Resource(id)
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDeployment_ResourceDetail(t *testing.T) {
	var d Deployment
	data, err := d.ResourceDetail(id, "TestDeployment_Create.bpmn20.xml")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDeployment_ResourceContent(t *testing.T) {
	var d Deployment
	data, err := d.ResourceContent(id, "TestDeployment_Create.bpmn20.xml")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}
