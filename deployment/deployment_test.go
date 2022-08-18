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
	flowablesdk.Setup("http://127.0.0.1:8067/flowable-rest/", "rest-admin", "123456", 10086)
	m.Run()
}

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
	data, err := d.Detail("4343f14b-1ddc-11ed-b722-f85ea0b3f3f9-x")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

// 获取详情
func TestDeployment_Create(t *testing.T) {
	var d Deployment
	req := CreateRequest{
		FileName: "TestDeployment_Create",
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
	err := d.Delete("8807b075-1e9c-11ed-ad65-0242ac170002")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("delete success")
}

func TestDeployment_Resource(t *testing.T) {
	var d Deployment
	data, err := d.Resource("2a24cec3-1e9c-11ed-ad65-0242ac170002")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDeployment_ResourceDetail(t *testing.T) {
	var d Deployment
	data, err := d.ResourceDetail("2a24cec3-1e9c-11ed-ad65-0242ac170002", "TestDeployment_Create")
	if err != nil {
		t.Error(err)
		return
	}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDeployment_ResourceContent(t *testing.T) {
	var d Deployment
	data, err := d.ResourceContent("2a24cec3-1e9c-11ed-ad65-0242ac170002", "TestDeployment_Create")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data)
}
