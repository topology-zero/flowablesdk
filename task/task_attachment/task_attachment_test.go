package task_attachment

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/attachment"
)

//go:embed _js3.png
var file []byte

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const (
	id           = "73243e27-b256-11ed-b3e2-38f3ab6b92c1"
	attachmentId = "c59829da-b351-11ed-b7d1-38f3ab6b92c1"
)

func TestList(t *testing.T) {
	data, err := List(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAddUrl(t *testing.T) {
	data, err := AddUrl(id, attachment.AddAttachment{
		Name:        "Simple attachment",
		Description: "Simple attachment description",
		Type:        "simpleType",
		ExternalUrl: "http://www.baidu.org",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestAdd(t *testing.T) {
	data, err := Add(id, attachment.AddAttachment{
		Name:        "Simple attachment1",
		Description: "Simple attachment description1",
		Type:        "simpleType1",
		File:        bytes.NewReader(file),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestDelete(t *testing.T) {
	err := Delete(id, attachmentId, "1")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestDetail(t *testing.T) {
	data, err := Detail(id, attachmentId)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(data)
}
